package howard

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

func (s *howardServer) GetIncomeStatement(ctx context.Context, request *stark.GetReportRequest) (*stark.GetIncomeStatementReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		tz     = getTimeZoneStr(request.GetTimeZone())
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()

		revenues []*stark.Income
		profits  []*stark.Income
		data     []struct {
			Date           time.Time `json:"date"`
			TopUpAmount    float64   `json:"top_up_amount"`
			WithdrawAmount float64   `json:"withdraw_amount"`
		}

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetIncomeStatement", zap.Any("request", request))

	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantIDEQ(request.GetMerchantId()),
			entpm.StatusEQ(int32(stark.Status_COMPLETED)),
		).
		Modify(func(s *sql.Selector) {
			topUpStm := fmt.Sprintf(
				"IF(%s = %d, COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) * %d, 0)",
				s.C(entpm.FieldType),
				int32(stark.PaymentType_TOPUP),
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
				_USDTExchangeRate,
			)
			withdrawStm := fmt.Sprintf(
				"IF(%s = %d, COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) * %d, 0)",
				s.C(entpm.FieldType),
				int32(stark.PaymentType_WITHDRAW),
				bankTbl.C(entbank.FieldAmount),
				telcoTbl.C(enttelco.FieldAmount),
				eWalletTbl.C(entwallet.FieldAmount),
				cryptoTbl.C(entcrypto.FieldAmount),
				_USDTExchangeRate,
			)
			dateStm := fmt.Sprintf("date_format(convert_tz(%s,'%s','%s'), '%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt), _defaultTz, tz)
			if tz == _defaultTz {
				dateStm = fmt.Sprintf("date_format(%s,'%%Y%%m%%d')", s.C(entpm.FieldUpdatedAt))
			}
			s.
				Select(
					sql.As(s.C(entpm.FieldUpdatedAt), "date"),
					sql.As(sql.Sum(topUpStm), "top_up_amount"),
					sql.As(sql.Sum(withdrawStm), "withdraw_amount"),
				).
				LeftJoin(bankTbl).On(s.C(entpm.FieldID), bankTbl.C(entbank.PaymentColumn)).
				LeftJoin(telcoTbl).On(s.C(entpm.FieldID), telcoTbl.C(enttelco.PaymentColumn)).
				LeftJoin(eWalletTbl).On(s.C(entpm.FieldID), eWalletTbl.C(entwallet.PaymentColumn)).
				LeftJoin(cryptoTbl).On(s.C(entpm.FieldID), cryptoTbl.C(entcrypto.PaymentColumn)).
				GroupBy(dateStm)
		}).
		Scan(ctx, &data)
	if err != nil {
		logger.Error("GetTotalAmount error", zap.Error(err))
	}

	for _, d := range data {
		revenues = append(revenues, &stark.Income{
			Date:   timestamppb.New(d.Date),
			Amount: d.TopUpAmount,
		})
		profits = append(profits, &stark.Income{
			Date:   timestamppb.New(d.Date),
			Amount: d.TopUpAmount - d.WithdrawAmount,
		})
	}

	return &stark.GetIncomeStatementReply{
		Revenues: revenues,
		Profits:  profits,
	}, nil
}
