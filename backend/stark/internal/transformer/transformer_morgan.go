package transformer

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

const (
	keepLen   = 4
	maskValue = "*"
)

func GetPaymentTelcoDetail(detail *ent.PaymentTelcoDetail) *stark.TelcoPaymentDetail {
	return &stark.TelcoPaymentDetail{
		Id:            detail.ID,
		SerialNumber:  mask(detail.SerialNumber, maskValue, keepLen),
		CardId:        mask(detail.CardID, maskValue, keepLen),
		ChargedAmount: detail.ChargedAmount,
		Amount:        detail.Amount,
		TelcoName:     groot.TelcoName(detail.TelcoName),
		CreatedAt:     timestamppb.New(detail.CreatedAt),
		UpdatedAt:     timestamppb.New(detail.UpdatedAt),
	}
}

func mask(data string, maskValue string, keepLen int) string {
	dataLen := len(data)
	if keepLen > dataLen {
		return data
	}
	maskLen := dataLen - keepLen
	keep := data[maskLen:]
	mask := strings.Repeat(maskValue, maskLen)

	return mask + keep
}
