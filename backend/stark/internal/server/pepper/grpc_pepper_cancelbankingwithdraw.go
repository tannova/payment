package pepper

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) CancelBankingWithdraw(ctx context.Context, request *stark.CancelBankingWithdrawRequest) (*stark.CancelBankingWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, mexIDStr, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get mex id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_T)
	expectedStatus := int32(stark.Status_CREATED)
	nextStatus := int32(stark.Status_CANCELED)

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.GetPaymentId()),
				epayment.TypeEQ(expectedType),
				epayment.MethodEQ(expectedMethod),
				epayment.StatusEQ(expectedStatus),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("failed to get banking withdraw payment by id", zap.Error(err))
			return err
		}

		caller := getMexIDAudit(mexIDStr)
		_, err = tx.Client().Payment.UpdateOne(payment).
			SetStatus(nextStatus).
			SetUpdatedBy(caller).
			SetUpdatedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to update banking withdraw payment status", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetStatus(nextStatus).
			SetNote("Payment is cancelled by merchant user").
			SetPayment(payment).
			Save(ctx)

		if err != nil {
			logger.Error("failed to create banking withdraw payment revision", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.CancelBankingWithdrawReply{}, nil
}
