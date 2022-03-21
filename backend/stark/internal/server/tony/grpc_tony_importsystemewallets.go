package tony

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"time"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	swallet "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
)

func (s *tonyServer) ImportSystemEWallets(ctx context.Context, request *stark.ImportSystemEWalletsRequest) (*stark.ImportSystemEWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	for _, r := range request.GetRecords() {
		query := s.entClient.SystemEWallet.
			Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetEWalletName(int32(r.AccountWalletName)).
			SetStatus(int32(stark.EWalletStatus_EWALLET_IN_ACTIVE)).
			SetMerchantID(r.GetMerchantId()).
			SetAccountPhoneNumber(r.GetAccountPhoneNumber()).
			SetAccountName(r.GetAccountName()).
			SetBalance(r.GetBalance()).
			SetDailyBalance(r.GetDailyBalance()).
			SetDailyBalanceLimit(r.GetDailyBalanceLimit()).
			SetDailyUsedAmount(r.GetDailyUsedAmount())

		if r.AccountId != 0 {
			query.SetID(r.AccountId)
		}
		query.
			OnConflict(
				sql.ConflictColumns(swallet.FieldID),
			).
			UpdateNewValues().
			Exec(ctx)
	}
	return &stark.ImportSystemEWalletsReply{}, nil
}
