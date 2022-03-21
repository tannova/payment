package howard

import (
	"context"
	"fmt"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	entpm "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	entbank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	entcrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	entwallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	enttelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
)

func (s *howardServer) GetTotalAmount(ctx context.Context, request *stark.GetTotalAmountRequest) (*stark.GetTotalAmountReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	var (
		logger = logging.Logger(ctx)
		tz     = getTimeZoneStr(request.GetTimeZone())
		from   = request.GetFromDate().AsTime()
		to     = request.GetToDate().AsTime()
		loc    = time.FixedZone("UTC "+getTimeZoneStr(request.GetTimeZone()), int(request.GetTimeZone()*60*60))

		tmp    = make(map[string]*stark.TotalAmountDetail)
		result []*stark.TotalAmountDetail
		data   []struct {
			MerchantID int64     `json:"merchant_id"`
			Date       time.Time `json:"date"`
			Amount     float64   `json:"amount"`
		}

		bankTbl    = sql.Table(entbank.Table).As("btbl")
		telcoTbl   = sql.Table(enttelco.Table).As("ttbl")
		eWalletTbl = sql.Table(entwallet.Table).As("etbl")
		cryptoTbl  = sql.Table(entcrypto.Table).As("ctbl")
	)
	if !request.GetFromDate().IsValid() || !request.GetToDate().IsValid() || from.After(to) {
		return nil, status.InvalidArgument("invalid from_date or to_date")
	}
	logger.Info("GetTotalAmount", zap.Any("request", request))

	// query
	err := s.entClient.Debug().Payment.Query().
		Where(
			entpm.UpdatedAtGTE(from),
			entpm.UpdatedAtLTE(to),
			entpm.MerchantIDIn(request.GetMerchants()...),
			entpm.Status(int32(stark.Status_COMPLETED)),
		).
		Modify(func(s *sql.Selector) {
			amountStm := fmt.Sprintf(
				"COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) + COALESCE(%s, 0) * %d",
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
	if err != nil {
		logger.Error("GetTotalAmount error", zap.Error(err))
	}

	// resulting
	for _, d := range data {
		key := d.Date.In(loc).Format(_dateFmt)
		if _, found := tmp[key]; !found {
			tmp[key] = &stark.TotalAmountDetail{Date: timestamppb.New(d.Date)}
		}
		tmp[key].MerchantAmounts = append(tmp[key].MerchantAmounts, &stark.TotalAmountMerchant{
			MerchantId: d.MerchantID,
			Amount:     uint64(d.Amount),
		})
	}

	for _, detail := range tmp {
		result = append(result, detail)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].GetDate().AsTime().Before(result[j].GetDate().AsTime()) })

	return &stark.GetTotalAmountReply{Details: result}, nil
}
