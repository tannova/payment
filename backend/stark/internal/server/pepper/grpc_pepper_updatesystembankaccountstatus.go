package pepper

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	sba "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) UpdateSystemBankAccountStatus(ctx context.Context, request *stark.UpdateSystemBankAccountStatusRequest) (*stark.UpdateSystemBankAccountStatusReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		acc, err := tx.Client().SystemBankAccount.
			Query().
			Where(
				sba.ID(request.Id),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("could not get the system bank account", zap.Error(err), zap.Int64("ID", request.Id))
			return status.Internal(err.Error())
		}
		_, err = tx.Client().SystemBankAccount.UpdateOne(acc).
			SetStatus(int32(request.Status)).
			SetUpdatedBy(tellerID).
			Save(ctx)
		if err != nil {
			logger.Error("could not update system bank account status", zap.Error(err))
			return status.Internal(err.Error())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &stark.UpdateSystemBankAccountStatusReply{}, nil
}
