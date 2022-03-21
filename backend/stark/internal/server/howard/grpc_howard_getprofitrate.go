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
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

type paymentMethodStats struct {
	PaymentType stark.PaymentType `json:"type"`
	BAmount     float64           `json:"b_amount"`
	EAmount     float64           `json:"e_amount"`
	TAmount     float64           `json:"t_amount"`
	CAmount     float64           `json:"c_amount"`
}

func (s *howardServer) GetProfitRate(ctx context.Context, request *stark.GetReportRequest) (*stark.GetProfitRateReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()

		totalWithdraw float64
		totalTopUp    float64
		profitRate    float64
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetProfitRate", zap.Any("request", request))

	wStats, err := s.getMethodAmountStats(ctx, from, to, request.GetMerchantId(), stark.PaymentType_WITHDRAW)
	if err != nil {
		logger.Error("failed to get paymentMethodStats", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	if wStats == nil {
		wStats = &paymentMethodStats{}
	}

	tStats, err := s.getMethodAmountStats(ctx, from, to, request.GetMerchantId(), stark.PaymentType_TOPUP)
	if err != nil {
		logger.Error("failed to get paymentMethodStats", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	if tStats == nil {
		tStats = &paymentMethodStats{}
	}

	totalTopUp = tStats.BAmount + tStats.EAmount + tStats.TAmount + tStats.CAmount
	totalWithdraw = wStats.BAmount + wStats.EAmount + wStats.TAmount + wStats.CAmount
	if totalTopUp > 0 {
		profitRate = (totalTopUp - totalWithdraw) / totalTopUp * 100.0
	}

	return &stark.GetProfitRateReply{
		TotalRevenue: totalTopUp,
		TotalProfit:  totalTopUp - totalWithdraw,
		ProfitRate:   profitRate,
	}, nil
}

func (s *howardServer) getMethodAmountStats(ctx context.Context, from, to time.Time, merchantID int64, pType stark.PaymentType) (*paymentMethodStats, error) {
	var (
		data []*paymentMethodStats

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)

	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantIDEQ(merchantID),
			entpm.StatusEQ(int32(stark.Status_COMPLETED)),
			entpm.Type(int32(pType)),
		).
		Modify(func(s *sql.Selector) {
			bAmountStm := fmt.Sprintf("COALESCE(%s, 0)", bankTbl.C(entbank.FieldAmount))
			eAmountStm := fmt.Sprintf("COALESCE(%s, 0)", eWalletTbl.C(entbank.FieldAmount))
			tAmountStm := fmt.Sprintf("COALESCE(%s, 0)", telcoTbl.C(entbank.FieldAmount))
			cAmountStm := fmt.Sprintf("COALESCE(%s, 0) * %d", cryptoTbl.C(entbank.FieldAmount), _USDTExchangeRate)

			s.
				Select(
					s.C(entpm.FieldType),
					sql.As(sql.Sum(bAmountStm), "b_amount"),
					sql.As(sql.Sum(eAmountStm), "e_amount"),
					sql.As(sql.Sum(tAmountStm), "t_amount"),
					sql.As(sql.Sum(cAmountStm), "c_amount"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn))
		}).
		Scan(ctx, &data)
	if err != nil || len(data) == 0 {
		return nil, err
	}

	return data[0], nil
}
