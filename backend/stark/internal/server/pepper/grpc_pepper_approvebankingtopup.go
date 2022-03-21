package pepper

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	esb "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) ApproveBankingTopUp(ctx context.Context, request *stark.ApproveBankingTopUpRequest) (*stark.ApproveBankingTopUpReply, error) {
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

	expectedType := int32(stark.PaymentType_TOPUP)
	expectedMethod := int32(stark.MethodType_T)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)
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
		paymentDetail, err := payment.QueryPaymentBankingDetail().Only(ctx)
		if err != nil {
			logger.Error("failed to get banking top up payment detail", zap.Error(err))
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

		// Start network call to MEX to inform
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
			s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextFailedStatus, err.Error())
			return err
		}

		payment, err = s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextSuccessStatus, request.Note)
		if err != nil {
			return err
		}

		paymentDetail, err = s.entClient.PaymentBankingDetail.
			UpdateOne(paymentDetail).
			SetImageURL(request.ImageUrl).
			SetTxID(request.TxId).
			Save(ctx)
		if err != nil {
			logger.Error("update payment detail err", zap.Error(err))
			return err
		}

		completeRequest, err := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			makePaymentReply.Balance,
		)

		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}
		err = s.jarvisClient.CompletePayment(
			ctx,
			reply.Merchant.WebhookUrl,
			completeRequest,
			s.completeTopUpCallback)
		if err != nil {
			logger.Error("complete payment err", zap.Error(err))
			return err
		}

		systemBank, err := s.entClient.SystemBankAccount.Query().
			Where(esb.And(
				esb.BankName(paymentDetail.SystemAccountBankName),
				esb.AccountNumber(paymentDetail.SystemAccountNumber),
				esb.AccountName(paymentDetail.SystemAccountName),
			)).
			First(ctx)
		if err != nil {
			logger.Error("failed to get system bank", zap.Error(err))
			return err
		}

		_, err = systemBank.Update().
			AddDailyUsedAmount(1).
			AddDailyBalance(paymentDetail.Amount).
			SetLastUpdatedBalance(time.Now().UTC()).
			Save(ctx)

		if err != nil {
			logger.Error("failed to update daily info of system bank", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.ApproveBankingTopUpReply{}, nil
}

func (s *pepperServer) completeTopUpCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_T,
			ExpectedStatus:    stark.Status_APPROVED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_APPROVE_FAILED,
		})
}
