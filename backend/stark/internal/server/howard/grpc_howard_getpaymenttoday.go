package howard

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
)

func (s *howardServer) GetPaymentToday(ctx context.Context, request *stark.GetReportRequest) (*stark.GetPaymentTodayReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	type Stats struct {
		TotalTopUp             int64 `json:"top_up"`
		TotalTopUpCompleted    int64 `json:"top_up_completed"`
		TotalWithdraw          int64 `json:"withdraw"`
		TotalWithdrawCompleted int64 `json:"withdraw_completed"`
	}
	var (
		logger = logging.Logger(ctx)
		loc    = time.FixedZone("UTC "+getTimeZoneStr(request.GetTimeZone()), int(request.GetTimeZone()*60*60))
		now    = time.Now().In(loc)
		from   = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		to     = from.AddDate(0, 0, 1)
		data   []*Stats
		stats  *Stats
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetPaymentToday", zap.Any("request", request))

	// query
	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLT(to),
			entpm.MerchantIDEQ(request.GetMerchantId()),
		).
		Modify(func(s *sql.Selector) {
			tmplTotal := "IF(%s = %d, 1, 0)"
			topUpStm := fmt.Sprintf(tmplTotal, s.C(entpm.FieldType), int32(stark.PaymentType_TOPUP))
			withdrawStm := fmt.Sprintf(tmplTotal, s.C(entpm.FieldType), int32(stark.PaymentType_WITHDRAW))

			tmplComplete := "IF(%s = %d and %s = %d, 1, 0)"
			topUpCompleteStm := fmt.Sprintf(
				tmplComplete,
				s.C(entpm.FieldType), int32(stark.PaymentType_TOPUP),
				s.C(entpm.FieldStatus), int32(stark.Status_COMPLETED),
			)
			withdrawCompleteStm := fmt.Sprintf(
				tmplComplete,
				s.C(entpm.FieldType), int32(stark.PaymentType_WITHDRAW),
				s.C(entpm.FieldStatus), int32(stark.Status_COMPLETED),
			)

			s.Select(
				sql.As(sql.Sum(topUpStm), "top_up"),
				sql.As(sql.Sum(withdrawStm), "withdraw"),
				sql.As(sql.Sum(topUpCompleteStm), "top_up_completed"),
				sql.As(sql.Sum(withdrawCompleteStm), "withdraw_completed"),
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

	return &stark.GetPaymentTodayReply{
		TopUpCompletion: &stark.TopUpCompletion{
			Total:     stats.TotalTopUp,
			Completed: stats.TotalTopUpCompleted,
		},
		WithdrawCompletion: &stark.WithdrawCompletion{
			Total:     stats.TotalWithdraw,
			Completed: stats.TotalWithdrawCompleted,
		},
	}, nil
}
