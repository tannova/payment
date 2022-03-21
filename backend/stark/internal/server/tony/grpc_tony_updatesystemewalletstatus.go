package tony

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func (s *tonyServer) UpdateSystemEWalletStatus(ctx context.Context, request *stark.UpdateSystemEWalletStatusRequest) (*stark.UpdateSystemEWalletStatusReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		wallet, err := tx.Client().SystemEWallet.
			Query().
			Where(
				esw.ID(request.Id),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("could not get ewallet", zap.Error(err), zap.Int64("ID", request.Id))
			return err
		}
		_, err = tx.Client().SystemEWallet.UpdateOne(wallet).
			SetStatus(int32(request.Status)).
			SetUpdatedBy(tellerID).
			Save(ctx)
		if err != nil {
			logger.Error("could not update system ewallet status", zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &stark.UpdateSystemEWalletStatusReply{}, nil
}
