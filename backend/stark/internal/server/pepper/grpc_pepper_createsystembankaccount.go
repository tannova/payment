package pepper

import (
	"context"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
)

func (s *pepperServer) CreateSystemBankAccount(ctx context.Context, request *stark.CreateSystemBankAccountRequest) (*stark.CreateSystemBankAccountReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	code := s.validateSystemBankAccount(ctx, request)
	if code != codes.Code_OK {
		return nil, status.Error(code)
	}

	return s.createSystemBankAccount.Create(ctx, request)
}
