package getpaymentinfobypaymentcode

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

type GetPaymentInfoByPaymentCode interface {
	Get(ctx context.Context, request *stark.GetPaymentInfoByPaymentCodeRequest) (*stark.GetPaymentInfoByPaymentCodeReply, error)
}

func New(
	setting setting.Setting,
) GetPaymentInfoByPaymentCode {

	return &getPaymentInfoByPaymentCode{
		setting: setting,
	}
}

type getPaymentInfoByPaymentCode struct {
	setting setting.Setting
}

func (g getPaymentInfoByPaymentCode) Get(ctx context.Context, request *stark.GetPaymentInfoByPaymentCodeRequest) (*stark.GetPaymentInfoByPaymentCodeReply, error) {
	// format : [Method] + [Service Provider] + [Merchant ID] + [UserPaymentID]
	// [Method] : 1 char
	// [Service Provider] : 1 char
	// [Merchant ID] : 2 digits
	// [UserPaymentID] : 8 digits
	paymentCode := request.Code
	paymentConf := g.setting.GetConfigPayment(ctx)
	max := paymentConf.Method + paymentConf.Provider + paymentConf.MerchantId + paymentConf.UserPaymentId
	if len(paymentCode) != int(max) {
		return nil, status.Error(codes.Code_PAYMENT_CODE_INVALID)
	}

	// reply
	reply := &stark.GetPaymentInfoByPaymentCodeReply{}

	// [Method]
	method, ok := stark.MethodType_value[paymentCode[:1]]
	if !ok {
		method = int32(stark.MethodType_METHOD_UNSPECIFIED)
	}

	// [Service Name]
	switch stark.MethodType(method) {
	case stark.MethodType_T:
		{
			mBankNames := g.setting.GetMapBankName()
			bankName, ok := stark.BankName_value[mBankNames[paymentCode[1:2]]]
			if !ok {
				bankName = int32(stark.BankName_BANK_UNSPECIFIED)
			}
			reply.Payload = &stark.GetPaymentInfoByPaymentCodeReply_BankName{BankName: stark.BankName(bankName)}
		}
	case stark.MethodType_E:
		{
			mEWalletName := g.setting.GetMapEWalletName()
			eWalletName, ok := stark.EWalletName_value[mEWalletName[paymentCode[1:2]]]
			if !ok {
				eWalletName = int32(stark.EWalletName_EWALLET_NAME_UNSPECIFIED)
			}
			reply.Payload = &stark.GetPaymentInfoByPaymentCodeReply_EWalletName{EWalletName: stark.EWalletName(eWalletName)}
		}
	}

	// [Merchant ID]
	mexStr := paymentCode[2:4]
	reply.MerchantId = utils.ConvertStrToInt(mexStr)

	// [UserPaymentID]
	userStr := paymentCode[4:12]
	reply.MerchantUserId = utils.ConvertStrToInt(userStr)

	reply.PaymentMethod = stark.MethodType(method)

	return reply, nil
}
