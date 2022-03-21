package pepper

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *pepperServer) ImportSystemBankAccount(ctx context.Context, request *stark.ImportSystemBankAccountRequest) (*stark.ImportSystemBankAccountReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)

	for _, createSystemBankAccountRequest := range request.GetRecords() {
		_, err := s.createSystemBankAccount.Create(ctx, createSystemBankAccountRequest)
		if err != nil {
			logger.Error("failed to create system account", zap.Error(err))
			return nil, status.Internal(err.Error())
		}
	}

	return &stark.ImportSystemBankAccountReply{}, nil
}
