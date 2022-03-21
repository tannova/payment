package report

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	eBanking "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	ecrypto "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	eWallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	etelco "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

const (
	_USDTExchangeRate = 23000
)

type Report interface {
	GetRevenue(ctx context.Context, fromDate time.Time, toDate time.Time, merchantID int64, pbPaymentTypes []stark.PaymentType, pbPaymentMethods []stark.MethodType) (float64, error)
}

func New(
	entClient *ent.Client,
) Report {

	return &report{
		entClient: entClient,
	}
}

type report struct {
	entClient *ent.Client
}

func (s *report) GetRevenue(ctx context.Context, fromDate time.Time, toDate time.Time, merchantID int64, pbPaymentTypes []stark.PaymentType, pbPaymentMethods []stark.MethodType) (float64, error) {
	logger := logging.Logger(ctx)

	var data []struct {
		MerchantID int64   `json:"merchant_id"`
		Amount     float64 `json:"amount"`
	}

	filters := []predicate.Payment{
		epayment.UpdatedAtGTE(fromDate),
		epayment.UpdatedAtLTE(toDate),
		epayment.MerchantIDEQ(merchantID),
		epayment.StatusEQ(int32(stark.Status_COMPLETED)),
	}

	if len(pbPaymentTypes) > 0 {
		paymentTypes := make([]int32, len(pbPaymentTypes))
		for i, v := range pbPaymentTypes {
			paymentTypes[i] = int32(v)
		}
		filters = append(filters, epayment.TypeIn(paymentTypes...))
	}

	if len(pbPaymentMethods) > 0 {
		methods := make([]int32, len(pbPaymentMethods))
		for i, v := range pbPaymentMethods {
			methods[i] = int32(v)
		}
		filters = append(filters, epayment.MethodIn(methods...))
	}

	bankingTable := sql.Table(eBanking.Table).As("t1")
	cryptoTable := sql.Table(ecrypto.Table).As("t2")
	telcoTable := sql.Table(etelco.Table).As("t3")
	eWalletTable := sql.Table(eWallet.Table).As("t4")
	err := s.entClient.Debug().Payment.Query().
		Where(epayment.And(
			filters...,
		)).
		GroupBy(epayment.FieldMerchantID).
		Aggregate(
			func(s *sql.Selector) string {
				appendSelectStr := fmt.Sprintf(
					"merchant_id, (COALESCE(%s, 0) + COALESCE(%s, 0) * %d + COALESCE(%s, 0) + COALESCE(%s, 0)) as amount",
					sql.Sum(bankingTable.C(eBanking.FieldAmount)),
					sql.Sum(cryptoTable.C(ecrypto.FieldAmount)),
					_USDTExchangeRate,
					sql.Sum(telcoTable.C(etelco.FieldAmount)),
					sql.Sum(eWalletTable.C(eWallet.FieldAmount)),
				)
				s.LeftJoin(bankingTable).On(s.C(epayment.FieldID), bankingTable.C(eBanking.PaymentColumn)).
					AppendSelect(appendSelectStr)
				return sql.As(sql.Sum(bankingTable.C(eBanking.FieldAmount)), "bank_amount")
			},
			func(s *sql.Selector) string {
				s.LeftJoin(cryptoTable).On(s.C(epayment.FieldID), cryptoTable.C(ecrypto.PaymentColumn))
				return sql.As(sql.Sum(cryptoTable.C(eBanking.FieldAmount)), "crypto_amount")
			},
			func(s *sql.Selector) string {
				s.LeftJoin(telcoTable).On(s.C(epayment.FieldID), telcoTable.C(etelco.PaymentColumn))
				return sql.As(sql.Sum(telcoTable.C(etelco.FieldAmount)), "telco_amount")
			},
			func(s *sql.Selector) string {
				s.LeftJoin(eWalletTable).On(s.C(epayment.FieldID), eWalletTable.C(eWallet.PaymentColumn))
				return sql.As(sql.Sum(eWalletTable.C(eWallet.FieldAmount)), "e_wallet_amount")
			},
		).
		Scan(ctx, &data)
	if err != nil {
		logger.Error("failed to calculate data", zap.Error(err))
		return 0, err
	}

	if len(data) <= 0 {
		logger.Error("data is empty", zap.Error(err))
		return 0, err
	}

	return data[0].Amount, nil
}
