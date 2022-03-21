package vision

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *visionServer) GetPaymentDetail(
	ctx context.Context,
	request *stark.GetPaymentDetailRequest,
) (*stark.GetPaymentDetailReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)

	payment, err := s.entClient.Payment.Get(ctx, request.GetId())
	if err != nil {
		logger.Error("get payment error", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	revisions, err := payment.QueryRevisions().All(ctx)
	if err != nil {
		logger.Error("failed to get payment revision", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	revisionPbs := make([]*stark.Revision, len(revisions))
	for i, r := range revisions {
		revisionPbs[i] = transformer.GetPaymentRevision(r)
	}

	return &stark.GetPaymentDetailReply{
		Payment:       transformer.GetPayment(payment),
		PaymentDetail: transformer.GetPaymentDetail(ctx, payment),
		Revisions:     revisionPbs,
	}, nil
}
