package ultron

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func (s *ultronServer) RejectCryptoWithdraw(
	ctx context.Context,
	request *stark.RejectCryptoWithdrawRequest,
) (*stark.RejectCryptoWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}
	now := time.Now().UTC()
	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_C)
	expectedStatuses := []int32{
		int32(stark.Status_CREATED),
		int32(stark.Status_APPROVED),
	}
	nextSuccessStatus := int32(stark.Status_REJECTING)

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				ep.IDEQ(request.GetPaymentId()),
				ep.TypeEQ(expectedType),
				ep.MethodEQ(expectedMethod),
				ep.StatusIn(expectedStatuses...),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("get payment error", zap.Error(err))
			return err
		}

		paymentDetail, err := payment.QueryPaymentCryptoDetail().Only(ctx)
		if err != nil {
			logger.Error("get payment detail error", zap.Error(err))
			return err
		}

		payment, err = tx.Client().Payment.UpdateOne(payment).
			SetStatus(nextSuccessStatus).
			SetUpdatedBy(tellerID).
			SetUpdatedAt(now).
			Save(ctx)
		if err != nil {
			logger.Error("could not update status payment", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(now).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetStatus(nextSuccessStatus).
			SetNote(request.GetNote()).
			SetPayment(payment).
			Save(ctx)

		if err != nil {
			logger.Error("create revision err", zap.Error(err))
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
			s.completeRejectCryptoWithdrawCallback)

		if err != nil {
			logger.Error("complete payment err", zap.Error(err))
			return err
		}

		if request.IsMerchantCall {
			s.blackWidowClient.NotifyRejectPayment(ctx, &natasha.NotifyRejectPaymentRequest{
				MerchantId: payment.MerchantID,
				Reason:     request.Note,
				PaymentId:  payment.ID,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.RejectCryptoWithdrawReply{}, nil
}

func (s *ultronServer) completeRejectCryptoWithdrawCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_WITHDRAW,
			ExpectedMethod:    stark.MethodType_C,
			ExpectedStatus:    stark.Status_REJECTING,
			NextSuccessStatus: stark.Status_REJECTED,
			NextFailedStatus:  stark.Status_REJECT_FAILED,
		})
}
