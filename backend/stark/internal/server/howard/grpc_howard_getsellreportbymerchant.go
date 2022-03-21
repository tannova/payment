package howard

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

func (s *howardServer) GetSellReportByMerchant(ctx context.Context, request *stark.GetSellReportByMerchantRequest) (*stark.GetSellReportByMerchantReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()

		topUps    = make([]*stark.SaleReportItem, 0)
		withdraws = make([]*stark.SaleReportItem, 0)
		data      []struct {
			MerchantID int64   `json:"merchant_id"`
			Amount     float64 `json:"amount"`
			Quantity   int64   `json:"quantity"`
			Discount   float64 `json:"discount"`
			Type       int32   `json:"type"`
		}

		filters    []predicate.Payment
		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	if len(_currencies[request.GetCurrency()]) == 0 {
		return nil, status.InvalidArgument("unsupported this currency")
	}
	logger.Info("GetSellReportByMerchant", zap.Any("request", request))

	if len(request.MerchantIds) > 0 {
		filters = append(filters, epayment.MerchantIDIn(request.MerchantIds...))
	}
	filters = append(
		filters,
		entpm.UpdatedAtGTE(from),
		entpm.UpdatedAtLTE(to),
		entpm.StatusEQ(int32(stark.Status_COMPLETED)),
		entpm.MethodIn(_currencies[request.GetCurrency()]...),
	)

	err := s.entClient.Debug().Payment.Query().
		Where(entpm.And(filters...)).
		Modify(func(s *sql.Selector) {
			amountStm := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0)",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
			)

			s.
				Select(
					s.C(entpm.FieldMerchantID),
					sql.As(sql.Sum(amountStm), "amount"),
					sql.As(sql.Count(s.C(entpm.FieldID)), "quantity"),
					s.C(entpm.FieldType),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn)).
				GroupBy(s.C(entpm.FieldMerchantID), s.C(entpm.FieldType))

		}).
		Scan(ctx, &data)
	if err != nil {
		logger.Error("failed to get data for banking", zap.Error(err))
		return nil, err
	}

	for _, d := range data {
		item := &stark.SaleReportItem{
			Key: &stark.SaleReportItem_MerchantId{
				MerchantId: d.MerchantID,
			},
			Quantity: d.Quantity,
			Amount:   d.Amount,
			Discount: d.Discount,
			Average:  d.Amount / float64(d.Quantity),
			Revenue:  d.Amount + d.Discount,
		}
		if stark.PaymentType(d.Type) == stark.PaymentType_TOPUP {
			topUps = append(topUps, item)
		} else if stark.PaymentType(d.Type) == stark.PaymentType_WITHDRAW {
			withdraws = append(withdraws, item)
		}
	}

	return &stark.GetSellReportByMerchantReply{
		TopUps:    topUps,
		Withdraws: withdraws,
	}, nil
}
