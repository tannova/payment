package blackwidow

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/payment"
)

func (s *blackWidowServer) GetMerchantBalance(ctx context.Context, request *natasha.GetMerchantBalanceRequest) (*natasha.GetMerchantBalanceReply, error) {
	mexID := request.GetMerchantId()
	reply := &natasha.GetMerchantBalanceReply{
		Balance:                     0,
		TotalMoneyIn:                0,
		TotalMoneyOut:               0,
		ReceiptVoucher:              nil,
		PaymentVoucher:              nil,
		BalanceForMexWithdrawProfit: 0,
		BalanceForMexWithdrawFunds:  0,
		Profit:                      0,
	}

	balances, err := s.repository.Payment.GetMerchantBalances(ctx, mexID, request.GetFromDate().AsTime(), request.GetToDate().AsTime())
	if err != nil {
		return reply, err
	}

	reply.Balance = balances.Balance
	reply.Profit = balances.Profit
	reply.TotalMoneyIn = balances.TotalMoneyIn
	reply.TotalMoneyOut = balances.TotalMoneyOut
	reply.BalanceForMexWithdrawProfit = balances.BalanceForMexWithdrawProfit
	reply.BalanceForMexWithdrawFunds = balances.BalanceForMexWithdrawFunds
	totalReceiptVoucherIn30Days := uint64(s.getTotalPaymentIn30DaysByType(
		ctx,
		mexID,
		int32(natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL),
		int32(natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION),
	))
	reply.ReceiptVoucher = &natasha.VoucherSummary{
		Total:          balances.TotalReceiptVoucher,
		TotalIn_30Days: totalReceiptVoucherIn30Days,
		Percent:        calcPercentage(totalReceiptVoucherIn30Days, balances.TotalReceiptVoucher),
	}
	totalPaymentVoucherIn30Days := uint64(-s.getTotalPaymentIn30DaysByType(
		ctx,
		mexID,
		int32(natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT),
		int32(natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS),
	))
	reply.PaymentVoucher = &natasha.VoucherSummary{
		Total:          balances.TotalPaymentVoucher,
		TotalIn_30Days: totalPaymentVoucherIn30Days,
		Percent:        calcPercentage(totalPaymentVoucherIn30Days, balances.TotalPaymentVoucher),
	}

	return reply, nil
}

func (s *blackWidowServer) getBalance(ctx context.Context, mexID int64) (result int64, err error) {
	var balance []struct {
		MerchantID int64 `json:"merchant_id"`
		Sum        int64 `json:"sum"`
	}

	err = s.eclient.Payment.
		Query().
		Where(
			epayment.MerchantIDEQ(mexID),
		).
		GroupBy(epayment.FieldMerchantID).
		Aggregate(ent.Sum(epayment.FieldAmount)).
		Scan(ctx, &balance)
	if err != nil {
		return result, err
	}
	// Consider no payment & balance is zero
	if len(balance) == 0 {
		return result, nil
	}

	return balance[0].Sum, nil
}

func (s *blackWidowServer) getTotalPaymentIn30DaysByType(ctx context.Context, mexID int64, types ...int32) int64 {
	fromDate := time.Now().UTC().AddDate(0, 0, -30)
	toDate := time.Now().UTC().AddDate(0, 0, 0)

	var data []struct {
		MerchantID int64 `json:"merchant_id"`
		Type       int32 `json:"type"`
		Sum        int64 `json:"sum"`
	}

	err := s.eclient.Payment.
		Query().
		Where(
			epayment.MerchantIDEQ(mexID),
			epayment.TypeIn(types...),
			epayment.CreatedAtGTE(fromDate),
			epayment.CreatedAtLTE(toDate),
		).
		GroupBy(epayment.FieldType, epayment.FieldMerchantID).
		Aggregate(ent.Sum(epayment.FieldAmount)).
		Scan(ctx, &data)

	if err == nil {
		if len(data) > 0 {
			return data[0].Sum
		}
	}

	s.logger.Warn(fmt.Sprintf("Error getTotalPaymentIn30DaysByType: %v", zap.Error(err)))
	return 0
}

func calcPercentage(a, b uint64) float64 {
	return math.Round(((float64(a)/float64(b))*100)*100) / 100
}

type SumBalance struct {
	MerchantID int64 `json:"merchant_id"`
	Type       int32 `json:"type"`
	Sum        int64 `json:"amount"`
}
