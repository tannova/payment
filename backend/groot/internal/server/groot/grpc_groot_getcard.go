package groot

import (
	"context"
	"go.uber.org/zap"

	ie "gitlab.com/mcuc/monorepo/backend/groot/internal/error"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/transformer"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
)

func (s *grootServer) GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	return s.processor.GetCard(ctx, request)
}

func (s *grootServer) getCardFromDB(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	s.logger.Debug("request", zap.Any("request", request))
	c, err := s.repository.Telco.GetAvailable(ctx, int64(request.TelcoName), request.Amount)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ie.NotFound
		}
		s.logger.Error("can not get card from db", zap.Error(err))
		return nil, ie.Internal(err.Error())
	}

	return transformer.GetTelco(c), nil
}
