package transformer

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func GetPayment(payment *ent.Payment) *stark.Payment {
	return &stark.Payment{
		Id:         payment.ID,
		MerchantId: payment.MerchantID,
		Method:     stark.MethodType(payment.Method),
		CreatedAt:  timestamppb.New(payment.CreatedAt),
		UpdatedAt:  timestamppb.New(payment.UpdatedAt),
		CreatedBy:  payment.CreatedBy,
		UpdatedBy:  payment.UpdatedBy,
		Type:       stark.PaymentType(payment.Type),
		Status:     stark.Status(payment.Status),
	}
}

func GetPaymentDetail(ctx context.Context, payment *ent.Payment) *stark.PaymentDetail {
	paymentDetail := &stark.PaymentDetail{}
	switch stark.MethodType(payment.Method) {
	case stark.MethodType_T:
		detail, _ := payment.QueryPaymentBankingDetail().Only(ctx)
		paymentDetail.Payload = &stark.PaymentDetail_Banking{
			Banking: GetPaymentBankingDetail(detail),
		}
	case stark.MethodType_E:
		detail, _ := payment.QueryPaymentEWalletDetail().Only(ctx)
		paymentDetail.Payload = &stark.PaymentDetail_EWallet{
			EWallet: GetPaymentEWalletDetail(detail),
		}
	case stark.MethodType_C:
		detail, _ := payment.QueryPaymentCryptoDetail().Only(ctx)
		paymentDetail.Payload = &stark.PaymentDetail_Crypto{
			Crypto: GetPaymentCryptoDetail(detail),
		}
	case stark.MethodType_P:
		detail, _ := payment.QueryPaymentTelcoDetail().Only(ctx)
		paymentDetail.Payload = &stark.PaymentDetail_Telco{
			Telco: GetPaymentTelcoDetail(detail),
		}
	}

	return paymentDetail
}

func GetPaymentRevision(revision *ent.Revision) *stark.Revision {
	return &stark.Revision{
		CreatedAt: timestamppb.New(revision.CreatedAt),
		CreatedBy: revision.CreatedBy,
		Status:    stark.Status(revision.Status),
		Note:      revision.Note,
	}
}
