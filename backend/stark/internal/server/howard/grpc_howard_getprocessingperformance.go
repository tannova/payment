package howard

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
)

func (s *howardServer) GetProcessingPerformance(ctx context.Context, request *stark.GetProcessingPerformanceRequest) (*stark.GetProcessingPerformanceReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	type Stats struct {
		Total          int64 `json:"total"`
		TotalCompleted int64 `json:"total_completed"`
		TotalFailed    int64 `json:"total_failed"`
		TotalWaiting   int64 `json:"total_waiting"`
	}
	var (
		logger = logging.Logger(ctx)
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()
		data   []*Stats
		stats  *Stats
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetProcessingPerformance", zap.Any("request", request))

	// query
	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantIDEQ(request.GetMerchantId()),
		).
		Modify(func(s *sql.Selector) {
			template := "IF(%s in (%s), 1, 0)"
			cStm := fmt.Sprintf(template, s.C(entpm.FieldStatus), sliceInt32ToString(_completedStatuses, ","))
			fStm := fmt.Sprintf(template, s.C(entpm.FieldStatus), sliceInt32ToString(_failedStatuses, ","))
			wStm := fmt.Sprintf(template, s.C(entpm.FieldStatus), sliceInt32ToString(_waitingStatuses, ","))

			s.Select(
				sql.As(sql.Count(s.C(entpm.FieldID)), "total"),
				sql.As(sql.Sum(cStm), "total_completed"),
				sql.As(sql.Sum(fStm), "total_failed"),
				sql.As(sql.Sum(wStm), "total_waiting"),
			)
		}).
		Scan(ctx, &data)
	if err != nil {
		logger.Error("getProcessingPerformance error", zap.Error(err))
	}
	if len(data) == 0 {
		stats = &Stats{}
	} else {
		stats = data[0]
	}

	// percent
	completedPercent := calculateProcessingPercent(stats.TotalCompleted, stats.Total)
	failedPercent := calculateProcessingPercent(stats.TotalFailed, stats.Total)
	waitingPercent := calculateProcessingPercent(stats.TotalWaiting, stats.Total)

	return &stark.GetProcessingPerformanceReply{
		MerchantId: request.GetMerchantId(),
		TotalOrder: uint64(stats.Total),
		Successfully: &stark.StatisticDetail{
			Number:  uint64(stats.TotalCompleted),
			Percent: completedPercent,
		},
		Failed: &stark.StatisticDetail{
			Number:  uint64(stats.TotalFailed),
			Percent: failedPercent,
		},
		Waiting: &stark.StatisticDetail{
			Number:  uint64(stats.TotalWaiting),
			Percent: waitingPercent,
		},
	}, nil
}
