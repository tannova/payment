package natasha

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

func (s *natashaServer) GetMerchant(ctx context.Context, request *natasha.GetMerchantRequest) (*natasha.GetMerchantReply, error) {
	if err := request.Validate(); err != nil {
		s.logger.Warn("invalid request", zap.Error(err))
		return nil, err
	}

	mex, err := s.repository.Merchant.GetMerchant(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &natasha.GetMerchantReply{
		Merchant: transformer.GetMerchant(mex),
	}, nil
}
