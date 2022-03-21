package tony

import (
	"context"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *tonyServer) CreateSystemEWallet(
	ctx context.Context,
	request *stark.CreateSystemEWalletRequest,
) (*stark.CreateSystemEWalletReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	code := s.validateCreateSystemEWallet(ctx, request)
	if code != codes.Code_OK {
		return nil, status.Error(code)
	}

	_, err = s.entClient.SystemEWallet.
		Create().
		SetCreatedAt(time.Now().UTC()).
		SetCreatedBy(tellerID).
		SetUpdatedBy(tellerID).
		SetEWalletName(int32(request.AccountWalletName)).
		SetStatus(int32(stark.EWalletStatus_EWALLET_IN_ACTIVE)).
		SetMerchantID(request.GetMerchantId()).
		SetAccountPhoneNumber(request.GetAccountPhoneNumber()).
		SetAccountName(request.GetAccountName()).
		SetBalance(request.GetBalance()).
		SetDailyBalance(request.GetDailyBalance()).
		SetDailyBalanceLimit(request.GetDailyBalanceLimit()).
		SetDailyUsedAmount(request.GetDailyUsedAmount()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &stark.CreateSystemEWalletReply{}, nil
}
