package pepper

import (
	"context"
	"fmt"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	eBankAccount "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"regexp"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/createsystembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/getpaymentinfobypaymentcode"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

const (
	_accountNumberMinLen        = 6
	_accountNameNonSpecialRegex = "^[a-zA-Z0-9_]+( [a-zA-Z0-9_]+)*$"
	_accountNumberOnlyNumber    = "^[0-9]+$"
)

func NewServer(service nightkit.Service,
	entClient *ent.Client,
	natashaCli natasha.NatashaClient,
	setting setting.Setting,
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode,
	blackWidowCli natasha.BlackWidowClient,
	natashaClient natasha.NatashaClient,
) stark.PepperServer {
	createSystemBankAccount := createsystembankaccount.New(entClient, natashaCli)
	return &pepperServer{
		logger:                      service.Logger(),
		stats:                       service.Stats(),
		entClient:                   entClient,
		natashaCli:                  natashaCli,
		createSystemBankAccount:     createSystemBankAccount,
		setting:                     setting,
		getPaymentInfoByPaymentCode: getPaymentInfoByPaymentCode,
		blackWidowCli:               blackWidowCli,
		jarvisClient:                jarvis.New(service.Logger()),
		natashaClient:               natashaClient,
	}
}

const (
	mexUserIDFmt   = "M%s"
	mexUserIDFmt64 = "M%d"
)

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

type pepperServer struct {
	stark.UnimplementedPepperServer
	logger                      *zap.Logger
	stats                       *statsd.Client
	entClient                   *ent.Client
	natashaCli                  natasha.NatashaClient
	setting                     setting.Setting
	createSystemBankAccount     createsystembankaccount.CreateSystemBankAccount
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode
	blackWidowCli               natasha.BlackWidowClient
	jarvisClient                jarvis.Client
	natashaClient               natasha.NatashaClient
}

func (s *pepperServer) updatePaymentStatus(ctx context.Context, tx tx.Tx, payment *ent.Payment, tellerStr string, status int32, note string) (*ent.Payment, error) {
	logger := logging.Logger(ctx)
	payment, err := tx.Client().Payment.UpdateOne(payment).
		SetStatus(status).
		SetUpdatedBy(tellerStr).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	if err != nil {
		logger.Error("failed to update banking top up payment status", zap.Error(err))
		return nil, err
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
		logger.Error("failed to create banking withdraw payment revision", zap.Error(err))
		return nil, err
	}
	return payment, nil
}

func (s *pepperServer) completeCallback(
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
	paymentDetail *ent.PaymentBankingDetail,
	currentBalance uint64,
) (*jarvis.CompletePaymentRequest, error) {
	return &jarvis.CompletePaymentRequest{
		PaymentId:     payment.ID,
		PaymentMethod: stark.MethodType(payment.Method),
		PaymentType:   stark.PaymentType(payment.Type),
		PaymentStatus: stark.Status(payment.Status),
		PaymentDetail: &jarvis.PaymentDetail{
			Banking: &stark.BankingPaymentDetail{
				Id:                        paymentDetail.ID,
				MerchantUserBankName:      stark.BankName(paymentDetail.MerchantUserBankName),
				MerchantUserAccountName:   paymentDetail.MerchantUserAccountName,
				MerchantUserAccountNumber: paymentDetail.MerchantUserAccountNumber,
				SystemBankName:            stark.BankName(paymentDetail.SystemAccountBankName),
				SystemAccountNumber:       paymentDetail.SystemAccountNumber,
				SystemAccountName:         paymentDetail.SystemAccountName,
				Amount:                    paymentDetail.Amount,
				ImageUrl:                  paymentDetail.ImageURL,
				TxId:                      paymentDetail.TxID,
				PaymentCode:               paymentDetail.PaymentCode,
				Fee:                       paymentDetail.Fee,
				MerchantUserId:            paymentDetail.MerchantUserID,
			},
		},
		MexCurrentBalance: currentBalance,
	}, nil
}

func (s *pepperServer) validateSystemBankAccount(ctx context.Context, request *stark.CreateSystemBankAccountRequest) codes.Code {
	logger := logging.Logger(ctx)

	s.logger.Info("Request", zap.Any("Data", request))
	match, err := regexp.MatchString(_accountNameNonSpecialRegex, request.AccountName)
	if err != nil {
		return codes.Code_INTERNAL
	}
	if !match {
		return codes.Code_BANK_ACCOUNT_NAME_INVALID
	}

	if len(request.AccountNumber) < _accountNumberMinLen {
		return codes.Code_BANK_ACCOUNT_NUMBER_INVALID
	}

	match, err = regexp.MatchString(_accountNumberOnlyNumber, request.AccountNumber)
	if err != nil {
		return codes.Code_INTERNAL
	}
	if !match {
		return codes.Code_BANK_ACCOUNT_NUMBER_INVALID
	}

	if request.Balance < 0 {
		return codes.Code_BANK_ACCOUNT_BALANCE_INVALID
	}

	if val, ok := stark.BankName_value[request.BankName]; !ok || val == 0 {
		return codes.Code_BANK_INVALID
	}
	if request.AccountId <= 0 {
		return codes.Code_BANK_INVALID
	}
	existed, err := s.entClient.SystemBankAccount.Query().
		Where(eBankAccount.And(
			eBankAccount.ID(request.AccountId),
		)).Exist(ctx)
	if err != nil {
		logger.Error("failed to existed account id", zap.Error(err))
		return codes.Code_INTERNAL
	}
	if existed {
		return codes.Code_BANK_ACCOUNT_ID_EXISTED
	}

	existed, err = s.entClient.SystemBankAccount.Query().
		Where(eBankAccount.And(
			eBankAccount.AccountNumber(request.AccountNumber),
		)).Exist(ctx)
	if err != nil {
		logger.Error("failed to existed account number", zap.Error(err))
		return codes.Code_INTERNAL
	}
	if existed {
		return codes.Code_BANK_ACCOUNT_EXISTED
	}

	return codes.Code_OK
}
