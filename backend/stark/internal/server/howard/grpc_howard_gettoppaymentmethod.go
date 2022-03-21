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
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

func (s *howardServer) GetTopPaymentMethod(ctx context.Context, request *stark.GetReportRequest) (*stark.GetTopPaymentMethodReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()
		data   []*stark.PaymentMethodRevenue

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetTopPaymentMethod", zap.Any("request", request))

	err := s.entClient.Debug().Payment.Query().
		Where(entpm.And(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantID(request.MerchantId),
			entpm.StatusEQ(int32(stark.Status_COMPLETED)),
		)).
		Modify(func(s *sql.Selector) {
			amountStm := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) * %d",
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
				_USDTExchangeRate,
			)

			s.
				Select(
					s.C(entpm.FieldType),
					s.C(entpm.FieldMethod),
					sql.As(sql.Sum(amountStm), "amount"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn)).
				GroupBy(s.C(entpm.FieldType), s.C(entpm.FieldMethod)).
				OrderBy(sql.Desc(sql.Sum(amountStm)))
		}).
		Scan(ctx, &data)
	if err != nil {
		logger.Error("failed to get data for banking", zap.Error(err))
		return nil, err
	}

	return &stark.GetTopPaymentMethodReply{TopPaymentMethodRevenue: data}, nil
}
