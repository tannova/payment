package pepper

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) RejectBankingWithdraw(ctx context.Context, request *stark.RejectBankingWithdrawRequest) (*stark.RejectBankingWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)

	_, tellerIDStr, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_T)
	expectedStatuses := []int32{
		int32(stark.Status_CREATED),
		int32(stark.Status_APPROVED),
	}
	nextStatus := int32(stark.Status_REJECTING)

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.GetPaymentId()),
				epayment.TypeEQ(expectedType),
				epayment.MethodEQ(expectedMethod),
				epayment.StatusIn(expectedStatuses...),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("failed to get banking withdraw payment by id", zap.Error(err))
			return err
		}

		paymentDetail, err := payment.QueryPaymentBankingDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment detail", zap.Error(err))
			return err
		}

		payment, err = tx.Client().Payment.UpdateOne(payment).
			SetStatus(nextStatus).
			SetUpdatedBy(tellerIDStr).
			SetUpdatedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to update banking withdraw payment status", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(tellerIDStr).
			SetUpdatedBy(tellerIDStr).
			SetStatus(nextStatus).
			SetNote(request.GetNote()).
			SetPayment(payment).
			Save(ctx)

		if err != nil {
			logger.Error("failed to create banking withdraw payment revision", zap.Error(err))
			return err
		}

		// Start network call to MEX to inform
		reply, err := s.natashaClient.GetMerchant(ctx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		completeRequest, err := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			// No need balance for this case
			// Will add this value when needed
			0,
		)
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}
		err = s.jarvisClient.CompletePayment(
			ctx,
			reply.Merchant.WebhookUrl,
			completeRequest,
			s.completeRejectWithdrawCallback)

		if request.IsMerchantCall {
			s.blackWidowCli.NotifyRejectPayment(ctx, &natasha.NotifyRejectPaymentRequest{
				MerchantId: payment.MerchantID,
				Reason:     request.Note,
				PaymentId:  payment.ID,
			})
		}

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.RejectBankingWithdrawReply{}, nil
}

func (s *pepperServer) completeRejectWithdrawCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_WITHDRAW,
			ExpectedMethod:    stark.MethodType_T,
			ExpectedStatus:    stark.Status_REJECTING,
			NextSuccessStatus: stark.Status_REJECTED,
			NextFailedStatus:  stark.Status_REJECT_FAILED,
		})
}
