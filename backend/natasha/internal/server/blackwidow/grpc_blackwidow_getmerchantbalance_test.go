package blackwidow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/service/merchant"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

const (
	testMexID = 191021
)

func Test_getSumBalances(t *testing.T) {
	tcs := []struct {
		payments    []*natasha.Payment
		expectation *merchant.MerchantBalances
	}{
		{
			payments: []*natasha.Payment{
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL,
					Amount:      100000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      1000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      2000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      20000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS,
					Amount:      -50000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION,
					Amount:      10000000,
				},
			},

			expectation: &merchant.MerchantBalances{
				Balance:                     77000000,
				TotalMoneyIn:                133000000,
				TotalMoneyOut:               56000000,
				BalanceForMexWithdrawProfit: 17000000,
				BalanceForMexWithdrawFunds:  50000000,
				TotalReceiptVoucher:         110000000,
				TotalPaymentVoucher:         50000000,
			},
		},

		{
			payments: []*natasha.Payment{
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL,
					Amount:      100000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      1000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      2000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      20000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS,
					Amount:      -50000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION,
					Amount:      10000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL,
					Amount:      100000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      30000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      10000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT,
					Amount:      -17000000,
				},
			},

			expectation: &merchant.MerchantBalances{
				Balance:                     197000000,
				TotalMoneyIn:                273000000,
				TotalMoneyOut:               76000000,
				BalanceForMexWithdrawProfit: 40000000,
				BalanceForMexWithdrawFunds:  150000000,
				TotalReceiptVoucher:         210000000,
				TotalPaymentVoucher:         67000000,
			},
		},

		{
			payments: []*natasha.Payment{
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL,
					Amount:      100000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      1000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      2000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      20000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS,
					Amount:      -50000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION,
					Amount:      10000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_WITHDRAW,
					Amount:      -3000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL,
					Amount:      100000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      30000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_USER_TOP_UP,
					Amount:      10000000,
				},
				{
					MerchantId:  testMexID,
					PaymentType: natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT,
					Amount:      -17000000,
				},
			},

			expectation: &merchant.MerchantBalances{
				Balance:                     194000000,
				TotalMoneyIn:                273000000,
				TotalMoneyOut:               79000000,
				BalanceForMexWithdrawProfit: 37000000,
				BalanceForMexWithdrawFunds:  150000000,
				TotalReceiptVoucher:         210000000,
				TotalPaymentVoucher:         70000000,
			},
		},
	}

	for idx, item := range tcs {
		tc := item
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()
			sumBalances := merchant.GetSumBalances(tc.payments)

			assert.Equal(t, sumBalances, tc.expectation)
		})
	}
}
