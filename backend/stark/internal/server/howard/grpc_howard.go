package howard

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"go.uber.org/zap"

	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/report"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

const (
	_USDTExchangeRate = 23000
	_defaultTz        = "+00:00"
	_dateFmt          = "02-01-2006"
)

var (
	_completedStatuses = []int32{
		int32(stark.Status_COMPLETED),
	}
	_failedStatuses = []int32{
		int32(stark.Status_REJECT_FAILED),
		int32(stark.Status_APPROVE_FAILED),
		int32(stark.Status_SUBMIT_FAILED),
	}
	_waitingStatuses = []int32{
		int32(stark.Status_CREATED),
		int32(stark.Status_CANCELED),
		int32(stark.Status_REJECTING),
		int32(stark.Status_REJECTED),
		int32(stark.Status_APPROVED),
		int32(stark.Status_SUBMITTED),
	}

	_currencies = map[stark.Currency][]int32{
		stark.Currency_VND: {
			int32(stark.MethodType_T),
			int32(stark.MethodType_E),
			int32(stark.MethodType_I),
			int32(stark.MethodType_P),
		},
		stark.Currency_USDT: {
			int32(stark.MethodType_C),
		},
	}
)

func NewServer(service nightkit.Service, entClient *ent.Client) stark.HowardServer {
	return &howardServer{
		logger:    service.Logger(),
		stats:     service.Stats(),
		entClient: entClient,
		report:    report.New(entClient),
	}
}

type howardServer struct {
	stark.UnimplementedHowardServer

	logger    *zap.Logger
	stats     *statsd.Client
	entClient *ent.Client
	report    report.Report
}

// utils

func curDateRangeByFilterType(filterType stark.StatisticFilterType, now time.Time) (time.Time, time.Time) {
	if filterType == stark.StatisticFilterType_STATISTIC_FILTER_MONTHLY {
		from := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return from, from.AddDate(0, 1, 0)
	}

	if filterType == stark.StatisticFilterType_STATISTIC_FILTER_WEEKLY {
		from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		for from.Weekday() != time.Monday {
			from = from.AddDate(0, 0, -1)
		}
		return from, from.AddDate(0, 0, 7)
	}

	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return from, from.AddDate(0, 0, 1)
}

func preDateRangeByFilterType(filterType stark.StatisticFilterType, now time.Time) (time.Time, time.Time) {
	if filterType == stark.StatisticFilterType_STATISTIC_FILTER_MONTHLY {
		to := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return to.AddDate(0, -1, 0), to
	}

	if filterType == stark.StatisticFilterType_STATISTIC_FILTER_WEEKLY {
		to := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		for to.Weekday() != time.Monday {
			to = to.AddDate(0, 0, -1)
		}
		return to.AddDate(0, 0, -7), to
	}

	to := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return to.AddDate(0, 0, -1), to
}

func calculateProcessingPercent(current, total int64) float64 {
	if total == 0 {
		return 0
	}

	return float64(current) / float64(total) * 100.0
}

func calculateStatisticPercent(curr, last float64) float64 {
	if last <= 0 {
		return 0
	}

	div := curr / last
	if div < 1 {
		return -1 * (1.0 - div) * 100
	} else {
		return (div - 1) * 100
	}
}

func getTimeZoneStr(timezone int32) string {
	if timezone < -12 || timezone > 14 {
		return _defaultTz
	}

	if timezone < 0 {
		return fmt.Sprintf("-%02d:00", -timezone)
	}

	return fmt.Sprintf("+%02d:00", timezone)
}

func sliceInt32ToString(lst []int32, sep string) string {
	var IDs []string
	for _, i := range lst {
		IDs = append(IDs, strconv.Itoa(int(i)))
	}

	return strings.Join(IDs, sep)
}
