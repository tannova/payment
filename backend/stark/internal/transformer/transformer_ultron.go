package transformer

import (
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func GetCryptoWallet(wallet *ent.CryptoWallet) *stark.CryptoWallet {
	return &stark.CryptoWallet{
		Id:                wallet.ID,
		Address:           wallet.Address,
		MerchantId:        wallet.MerchantID,
		CryptoType:        stark.CryptoType(wallet.CryptoType),
		CryptoNetworkType: stark.CryptoNetworkType(wallet.CryptoNetworkType),
		Status:            stark.CryptoWalletStatus(wallet.Status),
	}
}

func GetPaymentCryptoDetail(detail *ent.PaymentCryptoDetail) *stark.CryptoPaymentDetail {
	return &stark.CryptoPaymentDetail{
		CryptoType:        stark.CryptoType(detail.CryptoType),
		CryptoNetworkType: stark.CryptoNetworkType(detail.CryptoNetworkType),
		CryptoWalletName:  stark.CryptoWalletName(detail.CryptoWalletName),
		ReceiverAddress:   detail.ReceiverAddress,
		SenderAddress:     detail.SenderAddress,
		Amount:            detail.Amount,
		ReceivedAmount:    detail.ReceivedAmount,
		TxHash:            detail.TxHash,
		Fee:               detail.Fee,
		ImageUrl:          detail.ImageURL,
	}
}

func GetSystemCryptoHotWallet(wallet *ent.SystemCryptoHotWallet) *stark.SystemCryptoHotWallet {
	return &stark.SystemCryptoHotWallet{
		Id:                wallet.ID,
		Address:           wallet.Address,
		MerchantId:        wallet.MerchantID,
		CryptoType:        stark.CryptoType(wallet.CryptoType),
		CryptoNetworkType: stark.CryptoNetworkType(wallet.CryptoNetworkType),
		Balance:           wallet.Balance,
		TotalBalance:      wallet.TotalBalance,
		Status:            stark.CryptoHotWalletStatus(wallet.Status),
	}
}
