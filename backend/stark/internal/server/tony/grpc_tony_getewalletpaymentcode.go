package tony

import (
	"context"
	"fmt"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *tonyServer) GetEWalletPaymentCode(ctx context.Context, request *stark.GetEWalletPaymentCodeRequest) (*stark.GetEWalletPaymentCodeReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	mexID, _, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get mex id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	wallets, err := s.entClient.SystemEWallet.Query().
		Where(esw.And(
			esw.EWalletName(int32(request.EWalletName)),
			esw.Status(int32(stark.EWalletStatus_EWALLET_ACTIVE)),
			esw.MerchantID(mexID),
		)).
		Order(ent.Asc(
			esw.FieldDailyUsedAmount,
		)).
		All(ctx)
	if err != nil {
		logger.Error("failed to get system ewallet", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	var wallet *ent.SystemEWallet
	for _, w := range wallets {
		if w.DailyBalance >= w.DailyBalanceLimit {
			continue
		}
		wallet = w
		break
	}
	if wallet == nil {
		return nil, status.Error(codes.Code_EWALLET_NOT_FOUND_MERCHANT)
	}

	code := s.genEWalletPaymentCode(ctx, stark.MethodType_E, mexID, request.MerchantUserId, request.EWalletName)

	return &stark.GetEWalletPaymentCodeReply{
		Code:               code,
		EWalletName:        request.EWalletName,
		AccountPhoneNumber: wallet.AccountPhoneNumber,
		AccountName:        wallet.AccountName,
	}, nil
}

func (s *tonyServer) genEWalletPaymentCode(
	ctx context.Context,
	methodType stark.MethodType,
	mexID int64,
	mexUserID int64,
	eWalletName stark.EWalletName) string {
	merchantID := utils.ConvertIntToStr(
		mexID,
		int(s.setting.GetConfigPayment(ctx).MerchantId),
	)
	merchantUserID := utils.ConvertIntToStr(
		mexUserID,
		int(s.setting.GetConfigPayment(ctx).UserPaymentId),
	)
	shortName := ""
	for k, v := range s.setting.GetMapEWalletName() {
		if v == eWalletName.String() {
			shortName = k
		}
	}
	paymentCode := fmt.Sprintf("%s%s%s%s", methodType, shortName, merchantID, merchantUserID)

	return paymentCode
}
