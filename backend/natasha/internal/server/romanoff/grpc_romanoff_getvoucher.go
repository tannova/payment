package romanoff

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"go.uber.org/zap"
)

func (s *romanoffServer) GetVoucher(ctx context.Context, request *natasha.GetVoucherRequest) (*natasha.GetVoucherReply, error) {
	if err := request.Validate(); err != nil {
		s.logger.Warn("invalid request", zap.Error(err))
		return nil, err
	}

	result, err := s.eclient.Voucher.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &natasha.GetVoucherReply{
		Voucher: transformer.GetVoucher(result),
	}, nil
}
