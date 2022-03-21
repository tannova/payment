package pepper

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	code "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) VerifyMerchantUserBankAccount(ctx context.Context, request *stark.VerifyMerchantUserBankAccountRequest) (*stark.VerifyMerchantUserBankAccountReply, error) {
	if err := request.Validate(); err != nil {
		return nil, status.New(code.Code_MERCHANT_USER_BANK_ACCOUNT_INVALID, "invalid account number or account name")
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	_, err = s.entClient.MerchantUserBankAccount.Create().
		SetCreatedBy(tellerID).
		SetUpdatedBy(tellerID).
		SetBankName(int32(request.BankName)).
		SetAccountNumber(request.AccountNumber).
		SetAccountName(request.AccountName).
		Save(ctx)
	if err != nil {
		logger.Error("failed to verify merchant user bank account", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	return &stark.VerifyMerchantUserBankAccountReply{}, nil
}
