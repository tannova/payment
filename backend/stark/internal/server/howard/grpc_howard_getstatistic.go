package howard

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

type statisticStats struct {
	NumOfPayment int64   `json:"num_of_payment"`
	TotalAmount  float64 `json:"total_amount"`
	NumOfUser    int64   `json:"num_of_user"`
}

func (s *howardServer) GetStatistic(ctx context.Context, request *stark.GetStatisticRequest) (*stark.GetStatisticReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	var (
		err    error
		logger = logging.Logger(ctx)
		loc    = time.FixedZone("UTC "+getTimeZoneStr(request.GetTimeZone()), int(request.GetTimeZone()*60*60))
		now    = time.Now().In(loc)

		preStats, curStats *statisticStats
		preARPPU, curARPPU float64
		preStart, preEnd   = preDateRangeByFilterType(request.GetFilterType(), now)
		curStart, curEnd   = curDateRangeByFilterType(request.GetFilterType(), now)
	)
	logger.Info(
		"GetStatistic",
		zap.Any("request", request),
		zap.String("pre_range", fmt.Sprintf("%s - %s", preStart.Format(_dateFmt), preEnd.Format(_dateFmt))),
		zap.String("cur_range", fmt.Sprintf("%s - %s", curStart.Format(_dateFmt), curEnd.Format(_dateFmt))),
	)

	if preStats, err = s.getStatisticStats(ctx, preStart, preEnd, request.GetPaymentType()); err != nil {
		logger.Error("getStatisticStats error", zap.Error(err))
	}
	if curStats, err = s.getStatisticStats(ctx, curStart, curEnd, request.GetPaymentType()); err != nil {
		logger.Error("getStatisticStats error", zap.Error(err))
	}
	if preStats == nil {
		preStats = &statisticStats{}
	}
	if curStats == nil {
		curStats = &statisticStats{}
	}

	// arppu
	if preStats.NumOfUser > 0 {
		preARPPU = preStats.TotalAmount / float64(preStats.NumOfUser)
	}
	if curStats.NumOfUser > 0 {
		curARPPU = curStats.TotalAmount / float64(curStats.NumOfUser)
	}

	// percent
	paymentPercent := calculateStatisticPercent(float64(curStats.NumOfPayment), float64(preStats.NumOfPayment))
	amountPercent := calculateStatisticPercent(curStats.TotalAmount, preStats.TotalAmount)
	userPercent := calculateStatisticPercent(float64(curStats.NumOfUser), float64(preStats.NumOfUser))
	arppuPercent := calculateStatisticPercent(curARPPU, preARPPU)

	return &stark.GetStatisticReply{
		Order:  &stark.StatisticDetail{Number: uint64(curStats.NumOfPayment), Percent: paymentPercent},
		Amount: &stark.StatisticDetail{Number: uint64(curStats.TotalAmount), Percent: amountPercent},
		User:   &stark.StatisticDetail{Number: uint64(curStats.NumOfUser), Percent: userPercent},
		Arppu:  &stark.StatisticDetail{Number: uint64(curARPPU), Percent: arppuPercent},
	}, nil
}

func (s *howardServer) getStatisticStats(ctx context.Context, from, to time.Time, pmType stark.PaymentType) (*statisticStats, error) {
	var (
		tmp        []*statisticStats
		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)

	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLT(to),
			entpm.TypeEQ(int32(pmType)),
			entpm.StatusEQ(int32(stark.Status_COMPLETED)),
		).
		Modify(func(s *sql.Selector) {
			statement := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) * %d",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
				_USDTExchangeRate,
			)
			s.
				Select(
					sql.As(sql.Count(s.C(entpm.FieldID)), "num_of_payment"),
					sql.As(sql.Sum(statement), "total_amount"),
					sql.As(sql.Count(sql.Distinct(s.C(entpm.FieldMerchantUserID))), "num_of_user"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn))
		}).
		Scan(ctx, &tmp)

	if err != nil {
		return nil, err
	}
	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}
