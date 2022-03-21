package howard

import (
	"context"
	"fmt"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/enttest"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

const (
	//_driver = "sqlite3"
	//_url    = "file:ent?mode=memory&cache=shared&_fk=1"
	_driver = "mysql"
	_url    = "admin:THBGV3edzsGPvagU@tcp(localhost:53306)/db-stark?charset=utf8&parseTime=true"
)

func TestTimeZone(t *testing.T) {
	var (
		logger = logging.NewTmpLogger()
		loc    = time.FixedZone("CUSTOM", 7*60*60)
		now    = time.Now().In(loc)

		preStart, preEnd = preDateRangeByFilterType(stark.StatisticFilterType_STATISTIC_FILTER_WEEKLY, now)
		curStart, curEnd = curDateRangeByFilterType(stark.StatisticFilterType_STATISTIC_FILTER_WEEKLY, now)
	)

	logger.Info(
		"TestTimeZone",
		zap.String("now", fmt.Sprintf("%s - %d", now.String(), now.Unix())),
		zap.String("pre_range", fmt.Sprintf("%d - %d", preStart.Unix(), preEnd.Unix())),
		zap.String("cur_range", fmt.Sprintf("%d - %d", curStart.Unix(), curEnd.Unix())),
	)
}

func TestGetStatistic(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logging.NewTmpLogger()

		client   = enttest.Open(t, _driver, _url)
		from, to = preDateRangeByFilterType(stark.StatisticFilterType_STATISTIC_FILTER_WEEKLY, time.Now())
		data     []*statisticStats

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	logger.Info(
		"TestGetStatistic start",
		zap.String("from", from.Format(_dateFmt)),
		zap.String("to", to.Format(_dateFmt)),
	)

	err := client.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLT(to),
			entpm.TypeEQ(int32(stark.PaymentType_TOPUP)),
			entpm.StatusEQ(int32(stark.Status_COMPLETED)),
		).
		Modify(func(s *sql.Selector) {
			statement := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0)",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				fmt.Sprintf("%s*%d", cryptoTbl.C(entcrypto.FieldAmount), 23000),
			)
			s.
				Select(
					sql.As(sql.Count(s.C(entpm.FieldID)), "total"),
					sql.As(sql.Sum(statement), "total_amount"),
					sql.As(sql.Count(sql.Distinct(s.C(entpm.FieldMerchantUserID))), "total_user"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn))
		}).Scan(ctx, &data)

	logger.Info("TestGetStatistic end", zap.Any("data", data), zap.Error(err))
}

func TestGetProcessingPerformance(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logging.NewTmpLogger()

		client = enttest.Open(t, _driver, _url)
		to     = time.Now()
		from   = to.AddDate(0, 0, -7)
		data   []struct {
			Total          int64 `json:"total"`
			TotalCompleted int64 `json:"total_completed"`
			TotalFailed    int64 `json:"total_failed"`
			TotalWaiting   int64 `json:"total_waiting"`
		}
	)
	logger.Info("TestGetProcessingPerformance start")

	err := client.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLT(to),
		).
		Modify(func(s *sql.Selector) {
			//template := "case when %s in (%s) then 1 else 0 end"
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

	logger.Info("TestGetProcessingPerformance end", zap.Any("data", data), zap.Error(err))
}

func TestGetTotalAmount(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logging.NewTmpLogger()

		client      = enttest.Open(t, _driver, _url)
		tz          = getTimeZoneStr(7)
		to          = time.Now()
		from        = to.AddDate(0, 0, -7)
		merchantIDs = []int64{1, 5, 8}
		statuses    = []int32{int32(stark.Status_COMPLETED)}
		data        []struct {
			MerchantID int64     `json:"merchant_id"`
			Date       time.Time `json:"date"`
			Amount     float64   `json:"amount"`
		}

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	logger.Info(
		"TestGetTotalAmount start",
		zap.String("timezone", tz),
		zap.String("from", from.Format(_dateFmt)),
		zap.String("to", to.Format(_dateFmt)),
	)

	err := client.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantIDIn(merchantIDs...),
			entpm.StatusIn(statuses...),
		).
		Modify(func(s *sql.Selector) {
			amountStm := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0)",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				fmt.Sprintf("%s*%d", cryptoTbl.C(entcrypto.FieldAmount), _USDTExchangeRate),
			)
			dateStm := fmt.Sprintf("date_format(convert_tz(%s,'%s','%s'), '%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt), _defaultTz, tz)
			if tz == _defaultTz {
				dateStm = fmt.Sprintf("date_format(%s,'%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt))
			}

			s.
				Select(
					sql.As(s.C(entpm.FieldMerchantID), "merchant_id"),
					sql.As(s.C(entpm.FieldUpdatedAt), "date"),
					sql.As(sql.Sum(amountStm), "amount"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn)).
				GroupBy(s.C(entpm.FieldMerchantID), dateStm)
		}).Scan(ctx, &data)

	logger.Info("TestGetTotalAmount end", zap.Any("data", data), zap.Error(err))
}

func TestGetSaleReportByTimeRange(t *testing.T) {
	var (
		ctx    = context.Background()
		logger = logging.NewTmpLogger()

		client = enttest.Open(t, _driver, _url)
		tz     = getTimeZoneStr(7)
		to     = time.Now()
		from   = to.AddDate(0, 0, -7)

		data []struct {
			Date     time.Time `json:"date"`
			Amount   float64   `json:"amount"`
			Quantity int64     `json:"quantity"`
			Discount float64   `json:"discount"`
			Type     int32     `json:"type"`
		}

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	logger.Info(
		"TestGetSaleReportByTimeRange start",
		zap.String("timezone", tz),
		zap.String("from", from.Format(_dateFmt)),
		zap.String("to", to.Format(_dateFmt)),
	)

	err := client.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.Status(int32(stark.Status_COMPLETED)),
			entpm.MethodIn(_currencies[stark.Currency_VND]...),
		).
		Modify(func(s *sql.Selector) {
			amountStm := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0)",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
			)
			dateStm := fmt.Sprintf("date_format(convert_tz(%s,'%s','%s'), '%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt), _defaultTz, tz)
			if tz == _defaultTz {
				dateStm = fmt.Sprintf("date_format(%s,'%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt))
			}

			s.
				Select(
					sql.As(s.C(entpm.FieldUpdatedAt), "date"),
					sql.As(sql.Sum(amountStm), "amount"),
					sql.As(sql.Count(s.C(entpm.FieldID)), "quantity"),
					s.C(entpm.FieldType),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn)).
				GroupBy(dateStm, s.C(entpm.FieldType))
		}).
		Scan(ctx, &data)

	logger.Info("TestGetSaleReportByTimeRange end", zap.Any("data", data), zap.Error(err))
}
