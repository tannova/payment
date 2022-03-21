package transformer

import (
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetSystemBankAccount(account *ent.SystemBankAccount) *stark.SystemBank {
	return &stark.SystemBank{
		Id:                account.ID,
		AccountNumber:     account.AccountNumber,
		AccountName:       account.AccountName,
		BankName:          stark.BankName(account.BankName),
		Status:            stark.BankStatus(account.Status),
		MerchantId:        account.MerchantID,
		Branch:            account.Branch,
		Balance:           account.Balance,
		DailyBalanceLimit: account.DailyBalanceLimit,
		DailyBalance:      account.DailyBalance,
		DailyUsedAmount:   account.DailyUsedAmount,
	}
}

func GetPaymentBankingDetail(detail *ent.PaymentBankingDetail) *stark.BankingPaymentDetail {
	return &stark.BankingPaymentDetail{
		MerchantUserAccountNumber: detail.MerchantUserAccountNumber,
		MerchantUserBankName:      stark.BankName(detail.MerchantUserBankName),
		MerchantUserAccountName:   detail.MerchantUserAccountName,
		SystemAccountName:         detail.SystemAccountName,
		SystemAccountNumber:       detail.SystemAccountNumber,
		SystemBankName:            stark.BankName(detail.SystemAccountBankName),
		Amount:                    detail.Amount,
		Fee:                       detail.Fee,
		ImageUrl:                  detail.ImageURL,
		TxId:                      detail.TxID,
		MerchantUserId:            detail.MerchantUserID,
		PaymentCode:               detail.PaymentCode,
		CreatedAt:                 timestamppb.New(detail.CreatedAt),
		UpdatedAt:                 timestamppb.New(detail.UpdatedAt),
	}
}
