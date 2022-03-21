package morgan

import (
	"context"
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/telcopayment"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/telcosetting"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func NewServer(service nightkit.Service,
	entClient *ent.Client,
	natashaClient natasha.NatashaClient,
	blackWidowCli natasha.BlackWidowClient,
	grootCli groot.GrootClient,
) stark.MorganServer {
	jarvisClient := jarvis.New(service.Logger())
	return &morganServer{
		logger:        service.Logger(),
		stats:         service.Stats(),
		entClient:     entClient,
		jarvisClient:  jarvisClient,
		natashaClient: natashaClient,
		blackWidowCli: blackWidowCli,
		grootCli:      grootCli,
		telcoPayment:  telcopayment.New(entClient, natashaClient, jarvisClient, blackWidowCli, grootCli),
		telcoSetting:  telcosetting.New(entClient),
	}
}

const (
	mexUserIDFmt                 = "M%s"
	mexUserIDFmt64               = "M%d"
	_autoApprovalSettingName     = "top_up_auto_approval"
	_enableThirdPartySettingName = "enable_third_party"
)

func getMexIDAudit(mexID string) string {
	return fmt.Sprintf(mexUserIDFmt, mexID)
}

func getMexIDAudit64(mexID int64) string {
	return fmt.Sprintf(mexUserIDFmt64, mexID)
}

type morganServer struct {
	stark.UnimplementedMorganServer
	logger        *zap.Logger
	stats         *statsd.Client
	entClient     *ent.Client
	jarvisClient  jarvis.Client
	natashaClient natasha.NatashaClient
	blackWidowCli natasha.BlackWidowClient
	grootCli      groot.GrootClient
	telcoPayment  telcopayment.TelcoPayment
	telcoSetting  telcosetting.Setting
}

func (s *morganServer) updatePaymentStatus(ctx context.Context, tx tx.Tx, payment *ent.Payment, tellerStr string, status int32, note string) (*ent.Payment, error) {
	logger := logging.Logger(ctx)
	payment, err := tx.Client().Payment.UpdateOne(payment).
		SetStatus(status).
		SetUpdatedBy(tellerStr).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	if err != nil {
		logger.Error("failed to update telco top up payment status", zap.Error(err))
		return payment, err
	}

	_, err = tx.Client().Revision.Create().
		SetCreatedAt(time.Now().UTC()).
		SetCreatedBy(tellerStr).
		SetUpdatedBy(tellerStr).
		SetStatus(status).
		SetNote(note).
		SetPayment(payment).
		Save(ctx)

	if err != nil {
		logger.Error("failed to create telco withdraw payment revision", zap.Error(err))
		return payment, err
	}
	return payment, nil
}

func (s *morganServer) completeCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
	transition PaymentTransition,
) {
	txErr := tx.WithTransaction(context.Background(), s.logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				ep.IDEQ(request.PaymentId),
				ep.TypeEQ(int32(transition.ExpectedType)),
				ep.MethodEQ(int32(transition.ExpectedMethod)),
				ep.StatusEQ(int32(transition.ExpectedStatus)),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			s.logger.Error("get payment error", zap.Error(err))
			return err
		}
		caller := getMexIDAudit64(payment.MerchantID)

		paymentUpdate := tx.Client().Payment.
			UpdateOne(payment).
			SetUpdatedAt(time.Now().UTC())
		revisionCreate := tx.Client().Revision.Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetPaymentID(payment.ID).
			SetPayment(payment)

		if responseErr != nil {
			s.logger.Error("complete payment error", zap.Error(responseErr))
			paymentUpdate.SetStatus(int32(transition.NextFailedStatus))
			revisionCreate.SetStatus(int32(transition.NextFailedStatus))
			revisionCreate.SetNote(responseErr.Error())
		} else {
			paymentUpdate.SetStatus(int32(transition.NextSuccessStatus))
			revisionCreate.SetStatus(int32(transition.NextSuccessStatus))
			revisionCreate.SetNote(transition.GetMerchantSuccessNote())
		}

		_, err = paymentUpdate.Save(ctx)
		if err != nil {
			s.logger.Error("could not update status payment", zap.Error(err))
			return err
		}
		_, err = revisionCreate.Save(ctx)
		if err != nil {
			s.logger.Error("create revision err", zap.Error(err))
			return err
		}

		return nil
	})

	if txErr != nil {
		s.logger.Error("callback complete payment err", zap.Error(txErr))
	}
}

func createCompletePaymentRequest(
	ctx context.Context,
	payment *ent.Payment,
	paymentDetail *ent.PaymentTelcoDetail,
	currentBalance uint64,
) *jarvis.CompletePaymentRequest {
	return &jarvis.CompletePaymentRequest{
		PaymentId:     payment.ID,
		PaymentMethod: stark.MethodType(payment.Method),
		PaymentType:   stark.PaymentType(payment.Type),
		PaymentStatus: stark.Status(payment.Status),
		PaymentDetail: &jarvis.PaymentDetail{
			Telco: &stark.TelcoPaymentDetail{
				Id:            paymentDetail.ID,
				SerialNumber:  paymentDetail.SerialNumber,
				CardId:        paymentDetail.CardID,
				ChargedAmount: paymentDetail.ChargedAmount,
				Amount:        paymentDetail.Amount,
				TelcoName:     groot.TelcoName(paymentDetail.TelcoName),
			},
		},
		MexCurrentBalance: currentBalance,
	}
}

type PaymentTransition struct {
	ExpectedType      stark.PaymentType
	ExpectedMethod    stark.MethodType
	ExpectedStatus    stark.Status
	NextSuccessStatus stark.Status
	NextFailedStatus  stark.Status
}

func (p *PaymentTransition) GetMerchantSuccessNote() string {
	switch p.NextSuccessStatus {
	case stark.Status_CANCELED:
		return "Merchant canceled payment."
	case stark.Status_REJECTED:
		return "Merchant rejected payment."
	case stark.Status_APPROVED:
		return "Merchant approved payment."
	case stark.Status_SUBMITTED:
		return "Merchant submitted payment."
	case stark.Status_COMPLETED:
		return "Merchant completed payment."
	}

	return "Merchant succeeded on handling payment."
}

func (s *morganServer) completeRejectTopUpCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_REJECTING,
			NextSuccessStatus: stark.Status_REJECTED,
			NextFailedStatus:  stark.Status_REJECT_FAILED,
		})
}

func (s *morganServer) completeRejectWithdrawCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_WITHDRAW,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_REJECTING,
			NextSuccessStatus: stark.Status_REJECTED,
			NextFailedStatus:  stark.Status_REJECT_FAILED,
		})
}
