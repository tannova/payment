package morgan

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *morganServer) CompleteTelcoTopUp(ctx context.Context, request *stark.CompleteTelcoTopUpRequest) (*stark.CompleteTelcoTopUpReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)

	md, err := utils.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}
	outCtx := metadata.NewOutgoingContext(ctx, md)

	expectedType := int32(stark.PaymentType_TOPUP)
	expectedMethod := int32(stark.MethodType_P)
	expectedStatus := int32(stark.Status_APPROVED)
	nextFailedStatus := int32(stark.Status_APPROVE_FAILED)

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				epayment.IDEQ(request.GetPaymentId()),
				epayment.TypeEQ(expectedType),
				epayment.MethodEQ(expectedMethod),
				epayment.StatusEQ(expectedStatus),
			).
			WithPaymentBankingDetail().
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment by id", zap.Error(err))
			return err
		}
		paymentDetail, err := payment.QueryPaymentTelcoDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment detail", zap.Error(err))
			return err
		}

		paymentDetail.Update().
			SetChargedAmount(request.ChargedAmount).
			Save(ctx)

		caller := getMexIDAudit64(payment.MerchantID)

		reply, err := s.natashaClient.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		makePaymentReply, err := s.blackWidowCli.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     paymentDetail.Amount,
			Type:       natasha.PaymentType_USER_TOP_UP,
		})
		if err != nil {
			logger.Error("failed to make payment", zap.Error(err))
			s.updatePaymentStatus(ctx, tx, payment, caller, nextFailedStatus, err.Error())
			return err
		}

		completeRequest := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			makePaymentReply.Balance,
		)
		err = s.jarvisClient.CompletePayment(
			ctx,
			reply.Merchant.WebhookUrl,
			completeRequest,
			s.completeTopUpCallback)

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.CompleteTelcoTopUpReply{}, nil
}

func (s *morganServer) completeTopUpCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_P,
			ExpectedStatus:    stark.Status_APPROVED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_APPROVE_FAILED,
		})
}
