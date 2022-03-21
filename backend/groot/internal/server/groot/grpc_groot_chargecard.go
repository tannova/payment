package groot

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
)

func (s *grootServer) ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) (*groot.ChargeCardReply, error) {
	err := s.processor.ChargeCard(ctx, request)
	if err != nil {
		return nil, err
	}
	return &groot.ChargeCardReply{}, nil
}
