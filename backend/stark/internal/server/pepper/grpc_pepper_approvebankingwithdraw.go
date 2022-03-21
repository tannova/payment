package pepper

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) ApproveBankingWithdraw(ctx context.Context, request *stark.ApproveBankingWithdrawRequest) (*stark.ApproveBankingWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	_, tellerStr, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}
	metaData, err := utils.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}
	outCtx := metadata.NewOutgoingContext(ctx, metaData)

	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_T)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)
	//nextFailedStatus := int32(stark.Status_APPROVE_FAILED)

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
			logger.Error("failed to get banking top up payment by id", zap.Error(err))
			return err
		}

		paymentDetail, err := payment.QueryPaymentBankingDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking withdraw payment detail", zap.Error(err))
			return err
		}

		payment, err = payment.Update().
			SetApprovedBy(tellerStr).
			SetApprovedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to get update approve by", zap.Error(err))
			return err
		}

		payment, err = s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextSuccessStatus, request.Note)
		if err != nil {
			return err
		}

		// Start network call to MEX to inform
		getMerchantReply, err := s.natashaClient.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("failed to get merchant balance", zap.Error(err))
			return err
		}

		completeRequest, err := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			0,
		)
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		err = s.jarvisClient.CompletePayment(
			ctx,
			getMerchantReply.Merchant.WebhookUrl,
			completeRequest,
			s.completeWithdrawCallback)

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.ApproveBankingWithdrawReply{}, nil
}
