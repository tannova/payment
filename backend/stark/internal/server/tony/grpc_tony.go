package tony

import (
	"context"
	"fmt"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
	"regexp"
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"gitlab.com/greyhole/myid/pkg/extractor"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/getpaymentinfobypaymentcode"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func NewServer(
	service nightkit.Service,
	entClient *ent.Client,
	natashaClient natasha.NatashaClient,
	blackWidowClient natasha.BlackWidowClient,
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode,
	setting setting.Setting,
) stark.TonyServer {
	return &tonyServer{
		logger:                      service.Logger(),
		stats:                       service.Stats(),
		entClient:                   entClient,
		jarvisClient:                jarvis.New(service.Logger()),
		natashaClient:               natashaClient,
		blackWidowClient:            blackWidowClient,
		getPaymentInfoByPaymentCode: getPaymentInfoByPaymentCode,
		setting:                     setting,
	}
}

type tonyServer struct {
	stark.UnimplementedTonyServer

	logger                      *zap.Logger
	stats                       *statsd.Client
	entClient                   *ent.Client
	jarvisClient                jarvis.Client
	natashaClient               natasha.NatashaClient
	blackWidowClient            natasha.BlackWidowClient
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode
	setting                     setting.Setting
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

const (
	mexUserIDFmt   = "M%s"
	mexUserIDFmt64 = "M%d"
)

func getMexIDAudit(mexID string) string {
	return fmt.Sprintf(mexUserIDFmt, mexID)
}

func getMexIDAudit64(mexID int64) string {
	return fmt.Sprintf(mexUserIDFmt64, mexID)
}

func getUserID(ctx context.Context) (int64, string, error) {
	n, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return 0, "", err
	}
	return n, strconv.FormatInt(n, 10), nil
}

func (s *tonyServer) completeCallback(
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
	paymentDetail *ent.PaymentEWalletDetail,
	currentBalance uint64,
) (*jarvis.CompletePaymentRequest, error) {
	return &jarvis.CompletePaymentRequest{
		PaymentId:     payment.ID,
		PaymentMethod: stark.MethodType(payment.Method),
		PaymentType:   stark.PaymentType(payment.Type),
		PaymentStatus: stark.Status(payment.Status),
		PaymentDetail: &jarvis.PaymentDetail{
			EWallet: &stark.EWalletPaymentDetail{
				PaymentCode:                    paymentDetail.PaymentCode,
				EWalletName:                    stark.EWalletName(paymentDetail.EWalletName),
				MerchantUserAccountPhoneNumber: paymentDetail.MerchantUserAccountPhoneNumber,
				MerchantUserAccountName:        paymentDetail.MerchantUserAccountName,
				SystemAccountPhoneNumber:       paymentDetail.SystemAccountPhoneNumber,
				SystemAccountName:              paymentDetail.SystemAccountName,
				Amount:                         paymentDetail.Amount,
				Fee:                            paymentDetail.Fee,
				MerchantUserId:                 paymentDetail.MerchantUserID,
			},
		},
		MexCurrentBalance: currentBalance,
	}, nil
}

func (s *tonyServer) validateCreateSystemEWallet(
	ctx context.Context,
	request *stark.CreateSystemEWalletRequest,
) codes.Code {
	logger := logging.Logger(ctx)

	match, err := regexp.MatchString(_accountNameNonSpecialRegex, request.AccountName)
	if err != nil {
		return codes.Code_INTERNAL
	}
	if !match {
		return codes.Code_EWALLET_ACCOUNT_NAME_INVALID
	}

	if len(request.AccountPhoneNumber) < 10 || len(request.AccountPhoneNumber) > 11 {
		return codes.Code_EWALLET_PHONE_NUMBER_INVALID
	}

	match, err = regexp.MatchString(_accountNumberOnlyNumber, request.AccountPhoneNumber)
	if err != nil {
		return codes.Code_INTERNAL
	}
	if !match {
		return codes.Code_EWALLET_PHONE_NUMBER_INVALID
	}

	if request.Balance < 0 {
		return codes.Code_EWALLET_ACCOUNT_BALANCE_INVALID
	}

	if _, ok := stark.EWalletName_name[int32(request.AccountWalletName)]; !ok || request.GetAccountWalletName() == stark.EWalletName_EWALLET_NAME_UNSPECIFIED {
		return codes.Code_EWALLET_INVALID
	}
	if request.AccountId <= 0 {
		return codes.Code_EWALLET_INVALID
	}
	existed, err := s.entClient.SystemEWallet.Query().
		Where(esw.And(
			esw.ID(request.AccountId),
		)).Exist(ctx)
	if err != nil {
		logger.Error("failed to existed account id", zap.Error(err))
		return codes.Code_INTERNAL
	}
	if existed {
		return codes.Code_EWALLET_ACCOUNT_ID_EXISTED
	}

	existed, err = s.entClient.SystemEWallet.Query().
		Where(
			esw.EWalletNameEQ(int32(request.AccountWalletName)),
			esw.AccountPhoneNumberEQ(request.AccountPhoneNumber),
		).Exist(ctx)
	if err != nil {
		logger.Error("can not query wallet for validate", zap.Error(err))
		return codes.Code_INTERNAL
	}
	if existed {
		return codes.Code_EWALLET_ACCOUNT_EXISTED
	}

	return codes.Code_OK
}
