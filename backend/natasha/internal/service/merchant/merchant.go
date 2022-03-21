package merchant

import (
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

type MerchantBalances struct {
	Balance                     uint64
	TotalMoneyIn                uint64
	TotalMoneyOut               uint64
	BalanceForMexWithdrawProfit int64
	BalanceForMexWithdrawFunds  int64
	TotalReceiptVoucher         uint64
	TotalPaymentVoucher         uint64
	Profit                      int64
}

func GetSumBalances(payments []*natasha.Payment) *MerchantBalances {
	result := &MerchantBalances{
		Balance:                     0,
		BalanceForMexWithdrawProfit: 0,
		BalanceForMexWithdrawFunds:  0,
		TotalMoneyIn:                0,
		TotalMoneyOut:               0,
		TotalReceiptVoucher:         0,
		TotalPaymentVoucher:         0,
		Profit:                      0,
	}

	var charter int64 = 0
	var funds int64 = 0
	var profits int64 = 0
	var compens int64 = 0
	for _, payment := range payments {
		paymentType := natasha.PaymentType(payment.GetPaymentType())
		switch paymentType {
		case natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL:
			result.TotalMoneyIn += uint64(payment.GetAmount())
			result.TotalReceiptVoucher += uint64(payment.GetAmount())
			funds += payment.GetAmount()
			charter += payment.GetAmount()
		case natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION:
			result.TotalMoneyIn += uint64(payment.GetAmount())
			result.TotalReceiptVoucher += uint64(payment.GetAmount())
			compens += payment.GetAmount()
		case natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT:
			result.TotalMoneyOut += uint64(-payment.GetAmount())
			result.TotalPaymentVoucher += uint64(-payment.GetAmount())
			profits += payment.GetAmount() // negative
		case natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS:
			result.TotalMoneyOut += uint64(-payment.GetAmount())
			result.TotalPaymentVoucher += uint64(-payment.GetAmount())
			funds += payment.GetAmount() // negative
			charter += payment.GetAmount()
		case natasha.PaymentType_USER_TOP_UP:
			result.TotalMoneyIn += uint64(payment.GetAmount())
			if charter > funds {
				diff := charter - funds
				if diff < payment.GetAmount() {
					funds += diff
					profits += payment.GetAmount() - diff
				} else {
					funds += payment.GetAmount()
				}
			} else {
				profits += payment.GetAmount()
			}
		case natasha.PaymentType_USER_WITHDRAW:
			result.TotalMoneyOut += uint64(-payment.GetAmount())
			if compens > 0 {
				diff := compens + payment.GetAmount()
				if diff > 0 {
					compens += payment.GetAmount()
				} else {
					// diff is negative
					compens = 0
					if profits > 0 {
						diff += profits
						if diff > 0 {
							profits = diff
						} else {
							profits = 0
							funds += diff
						}
					} else {
						funds += diff
					}
				}
			} else if profits > 0 {
				diff := profits + payment.GetAmount()
				if diff > 0 {
					profits += payment.GetAmount()
				} else if diff <= 0 {
					profits = 0
					funds += diff
				}
			} else if funds > 0 {
				funds += payment.GetAmount()
			}
		}
	}

	balance := int64(result.TotalMoneyIn) - int64(result.TotalMoneyOut)
	if balance < 0 {
		balance = 0
	}
	result.Balance = uint64(balance)
	result.BalanceForMexWithdrawProfit = profits
	result.BalanceForMexWithdrawFunds = funds

	if result.TotalMoneyIn > result.TotalMoneyOut {
		result.Profit = int64(result.TotalMoneyIn - result.TotalMoneyOut)
	}

	return result
}
