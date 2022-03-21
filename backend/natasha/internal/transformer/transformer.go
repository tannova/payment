package transformer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

func GetMerchant(mex *ent.Merchant) *natasha.Merchant {
	return &natasha.Merchant{
		Id:           mex.ID,
		Name:         mex.Name,
		LogoPath:     mex.LogoPath,
		EmailContact: mex.EmailContact,
		WebhookUrl:   mex.WebhookURL,
		CreatedBy:    mex.CreatedBy,
		UpdatedBy:    mex.UpdatedBy,
		CreatedAt:    timestamppb.New(mex.CreatedAt),
		UpdatedAt:    timestamppb.New(mex.UpdatedAt),
	}
}

func GetPayment(payment *ent.Payment) *natasha.Payment {
	return &natasha.Payment{
		Id:          payment.ID,
		MerchantId:  payment.MerchantID,
		PaymentType: natasha.PaymentType(payment.Type),
		Amount:      payment.Amount,
		CreatedBy:   payment.CreatedBy,
		UpdatedBy:   payment.UpdatedBy,
		CreatedAt:   timestamppb.New(payment.CreatedAt),
		UpdatedAt:   timestamppb.New(payment.UpdatedAt),
	}
}

func GetVoucher(voucher *ent.Voucher) *natasha.Voucher {
	return &natasha.Voucher{
		Id:            voucher.ID,
		MerchantId:    voucher.MerchantID,
		PaymentId:     voucher.PaymentID,
		Amount:        voucher.Amount,
		Type:          natasha.PaymentType(voucher.Type),
		Status:        natasha.VoucherStatus(voucher.Status),
		ImageUrl:      voucher.ImageURL,
		Note:          voucher.CreatorNote,
		PayeeProvider: uint64(voucher.PayeeProvider),
		PayeeAccount:  voucher.PayeeAccount,
		PayeeName:     voucher.PayeeName,
		PayerProvider: uint64(voucher.PayerProvider),
		PayerAccount:  voucher.PayerAccount,
		PayerName:     voucher.PayerName,
		TxId:          voucher.TxID,
		HandledBy:     voucher.HandledBy,
		HandlerNote:   voucher.HandlerNote,
		CreatedBy:     voucher.CreatedBy,
		UpdatedBy:     voucher.UpdatedBy,
		CreatedAt:     timestamppb.New(voucher.CreatedAt),
		UpdatedAt:     timestamppb.New(voucher.UpdatedAt),
	}
}
