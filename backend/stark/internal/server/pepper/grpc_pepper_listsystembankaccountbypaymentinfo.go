package pepper

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	sba "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) ListSystemBankAccountByPaymentInfo(ctx context.Context, request *stark.ListSystemBankAccountByPaymentInfoRequest) (*stark.ListSystemBankAccountByPaymentInfoReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, _, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	banks, err := s.entClient.SystemBankAccount.Query().
		Where(sba.And(
			sba.MerchantID(request.MerchantId),
			sba.BankName(int32(request.BankName)),
		)).All(ctx)
	if err != nil {
		logger.Error("failed to query system bank account for payment", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	records := make([]*stark.SystemBank, len(banks))
	for i, b := range banks {
		records[i] = transformer.GetSystemBankAccount(b)
	}

	return &stark.ListSystemBankAccountByPaymentInfoReply{
		Records: records,
	}, nil
}
