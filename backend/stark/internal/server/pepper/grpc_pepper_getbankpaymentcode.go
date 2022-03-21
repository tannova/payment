package pepper

import (
	"context"
	"fmt"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	sba "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) GetBankPaymentCode(ctx context.Context, request *stark.GetBankPaymentCodeRequest) (*stark.GetBankPaymentCodeReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	mexID, _, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get mex id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	banks, err := s.entClient.SystemBankAccount.Query().
		Where(sba.And(
			sba.BankName(int32(request.BankName)),
			sba.Status(int32(stark.BankStatus_BANK_STATUS_ACTIVE)),
			sba.MerchantID(mexID),
		)).
		Order(ent.Asc(
			sba.FieldDailyUsedAmount,
		)).
		All(ctx)
	if err != nil {
		logger.Error("failed to get system bank account", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	var bank *ent.SystemBankAccount
	for _, b := range banks {
		if b.DailyBalance >= b.DailyBalanceLimit {
			continue
		}
		bank = b
		break
	}
	if bank == nil {
		return nil, status.Error(codes.Code_BANK_NOT_FOUND_MERCHANT)
	}

	code := s.genBankPaymentCode(ctx, stark.MethodType_T, mexID, request.MerchantUserId, request.BankName)

	return &stark.GetBankPaymentCodeReply{
		Code:          code,
		AccountName:   bank.AccountName,
		AccountNumber: bank.AccountNumber,
		Branch:        bank.Branch,
	}, nil
}

func (s *pepperServer) genBankPaymentCode(ctx context.Context, methodType stark.MethodType, mexID, mexUserID int64, bankName stark.BankName) string {
	merchantID := utils.ConvertIntToStr(mexID, int(s.setting.GetConfigPayment(ctx).MerchantId))
	merchantUserID := utils.ConvertIntToStr(mexUserID, int(s.setting.GetConfigPayment(ctx).UserPaymentId))
	shortName := ""
	for k, v := range s.setting.GetMapBankName() {
		if v == bankName.String() {
			shortName = k
		}
	}
	paymentCode := fmt.Sprintf("%s%s%s%s", methodType, shortName, merchantID, merchantUserID)

	return paymentCode
}
