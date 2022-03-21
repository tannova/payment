package howard

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *howardServer) GetAllocationTopUpRate(ctx context.Context, request *stark.GetReportRequest) (*stark.GetAllocationTopUpRateReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetAllocationTopUpRate", zap.Any("request", request))

	stats, err := s.getMethodAmountStats(ctx, from, to, request.GetMerchantId(), stark.PaymentType_TOPUP)
	if err != nil {
		logger.Error("failed to get paymentMethodStats", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	if stats == nil {
		stats = &paymentMethodStats{}
	}

	total := stats.BAmount + stats.EAmount + stats.TAmount + stats.CAmount
	bRate := 0.0
	eRate := 0.0
	tRate := 0.0
	cRate := 0.0
	if total != 0 {
		bRate = (stats.BAmount / total) * 100.0
		eRate = (stats.EAmount / total) * 100.0
		tRate = (stats.TAmount / total) * 100.0
		cRate = 100.0 - (bRate + eRate + tRate)
	}

	return &stark.GetAllocationTopUpRateReply{
		TotalTopup: total,
		TopUpAllocationRate: []*stark.AllocationDetail{
			{Method: stark.MethodType_T, Amount: stats.BAmount, Percent: bRate},
			{Method: stark.MethodType_E, Amount: stats.EAmount, Percent: eRate},
			{Method: stark.MethodType_P, Amount: stats.TAmount, Percent: tRate},
			{Method: stark.MethodType_C, Amount: stats.CAmount, Percent: cRate},
		},
	}, nil
}
