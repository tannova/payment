package pepper

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	em "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/merchantuserbankaccount"
)

func (s *pepperServer) GetIsMerchantUserBankAccountVerified(ctx context.Context, request *stark.GetIsMerchantUserBankAccountVerifiedRequest) (*stark.GetIsMerchantUserBankAccountVerifiedReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)

	verified, err := s.entClient.MerchantUserBankAccount.Query().Where(em.And(
		em.AccountName(request.AccountName),
		em.AccountNumber(request.AccountNumber),
	)).Exist(ctx)
	if err != nil {
		logger.Error("failed to get is merchant user bank account verified", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	return &stark.GetIsMerchantUserBankAccountVerifiedReply{
		Verified: verified,
	}, nil
}
