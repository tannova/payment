package transformer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func GetSystemEWallet(e *ent.SystemEWallet) *stark.SystemEWallet {
	return &stark.SystemEWallet{
		Id:                 e.ID,
		CreatedAt:          timestamppb.New(e.CreatedAt),
		UpdatedAt:          timestamppb.New(e.UpdatedAt),
		CreatedBy:          e.CreatedBy,
		UpdatedBy:          e.UpdatedBy,
		EWalletName:        stark.EWalletName(e.EWalletName),
		Status:             stark.EWalletStatus(e.Status),
		MerchantId:         e.MerchantID,
		AccountPhoneNumber: e.AccountPhoneNumber,
		AccountName:        e.AccountName,
		Balance:            e.Balance,
		DailyBalance:       e.DailyBalance,
		DailyBalanceLimit:  e.DailyBalanceLimit,
		DailyUsedAmount:    e.DailyUsedAmount,
	}
}

func GetSystemEWallets(se []*ent.SystemEWallet) []*stark.SystemEWallet {
	ews := []*stark.SystemEWallet{}
	for _, ew := range se {
		ews = append(ews, GetSystemEWallet(ew))
	}

	return ews
}

func GetPaymentEWalletDetail(detail *ent.PaymentEWalletDetail) *stark.EWalletPaymentDetail {
	return &stark.EWalletPaymentDetail{
		PaymentCode:                    detail.PaymentCode,
		EWalletName:                    stark.EWalletName(detail.EWalletName),
		MerchantUserId:                 detail.MerchantUserID,
		MerchantUserAccountName:        detail.MerchantUserAccountName,
		MerchantUserAccountPhoneNumber: detail.MerchantUserAccountPhoneNumber,
		SystemAccountName:              detail.SystemAccountName,
		SystemAccountPhoneNumber:       detail.SystemAccountPhoneNumber,
		Amount:                         detail.Amount,
		Fee:                            detail.Fee,
	}
}
