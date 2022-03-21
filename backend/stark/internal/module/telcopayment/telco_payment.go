package telcopayment

import (
	"context"
	"fmt"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/telcosetting"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"time"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

const (
	mexUserIDFmt   = "M%s"
	mexUserIDFmt64 = "M%d"
)

type TelcoPayment interface {
	ApproveTopUp(ctx context.Context, callerID string, request *stark.ApproveTelcoTopUpRequest) (*stark.ApproveTelcoTopUpReply, error)
}

func New(
	entClient *ent.Client,
	natashaCli natasha.NatashaClient,
	jarvisClient jarvis.Client,
	blackWidowCli natasha.BlackWidowClient,
	grootCli groot.GrootClient,
) TelcoPayment {
	return &telcoPayment{
		entClient:     entClient,
		natashaCli:    natashaCli,
		jarvisClient:  jarvisClient,
		blackWidowCli: blackWidowCli,
		grootCli:      grootCli,
		telcoSetting:  telcosetting.New(entClient),
	}
}

type telcoPayment struct {
	entClient     *ent.Client
	natashaCli    natasha.NatashaClient
	jarvisClient  jarvis.Client
	blackWidowCli natasha.BlackWidowClient
	grootCli      groot.GrootClient
	telcoSetting  telcosetting.Setting
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

func getMexIDAudit(mexID string) string {
	return fmt.Sprintf(mexUserIDFmt, mexID)
}

func getMexIDAudit64(mexID int64) string {
	return fmt.Sprintf(mexUserIDFmt64, mexID)
}

func createCompletePaymentRequest(
	ctx context.Context,
	payment *ent.Payment,
	paymentDetail *ent.PaymentTelcoDetail,
	currentBalance uint64,
) (*jarvis.CompletePaymentRequest, error) {
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
	}, nil
}

func (a telcoPayment) updatePaymentStatus(ctx context.Context, tx tx.Tx, payment *ent.Payment, tellerStr string, status int32, note string) (*ent.Payment, error) {
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

func (a telcoPayment) completeCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
	transition PaymentTransition,
) {
	logger := logging.Logger(nil)
	txErr := tx.WithTransaction(context.Background(), logger, a.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.PaymentId),
				epayment.TypeEQ(int32(transition.ExpectedType)),
				epayment.MethodEQ(int32(transition.ExpectedMethod)),
				epayment.StatusEQ(int32(transition.ExpectedStatus)),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("get payment error", zap.Error(err))
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
			logger.Error("complete payment error", zap.Error(responseErr))
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
			logger.Error("could not update status payment", zap.Error(err))
			return err
		}
		_, err = revisionCreate.Save(ctx)
		if err != nil {
			logger.Error("create revision err", zap.Error(err))
			return err
		}

		return nil
	})

	if txErr != nil {
		logger.Error("callback complete payment err", zap.Error(txErr))
	}
}
