package pepper

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

func (s *pepperServer) SubmitBankingWithdraw(ctx context.Context, request *stark.SubmitBankingWithdrawRequest) (*stark.SubmitBankingWithdrawReply, error) {
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
	expectedStatus := int32(stark.Status_APPROVED)
	nextSuccessStatus := int32(stark.Status_SUBMITTED)
	//nextFailedStatus := int32(stark.Status_SUBMIT_FAILED)

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

		paymentDetail, err = paymentDetail.Update().
			SetSystemAccountBankName(int32(request.BankName)).
			SetSystemAccountNumber(request.AccountNumber).
			SetSystemAccountName(request.AccountName).
			SetAmount(request.Amount).
			SetImageURL(request.ImageUrl).
			SetTxID(request.TxId).
			SetFee(request.Fee).
			Save(ctx)
		if err != nil {
			logger.Error("failed to update payment detail", zap.Error(err))
			return err
		}

		_, err = s.blackWidowCli.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     paymentDetail.Amount,
			Type:       natasha.PaymentType_USER_WITHDRAW,
		})
		if err != nil {
			logger.Error("failed to make payment", zap.Error(err))
			return err
		}

		// Start network call to MEX to inform
		reply, err := s.natashaClient.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		merchantBalanceReply, err := s.blackWidowCli.GetMerchantBalance(outCtx, &natasha.GetMerchantBalanceRequest{
			MerchantId: payment.MerchantID,
		})
		if err != nil {
			logger.Error("failed to get merchant balance", zap.Error(err))
			return err
		}

		payment, err = s.updatePaymentStatus(ctx, tx, payment, tellerStr, nextSuccessStatus, request.Note)
		if err != nil {
			logger.Error("failed to get update payment status", zap.Error(err))
			return err
		}

		completeRequest, err := createCompletePaymentRequest(
			ctx,
			payment,
			paymentDetail,
			merchantBalanceReply.Balance,
		)
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}

		err = s.jarvisClient.CompletePayment(
			ctx,
			reply.Merchant.WebhookUrl,
			completeRequest,
			s.completeWithdrawCallback)

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return &stark.SubmitBankingWithdrawReply{}, nil
}

func (s *pepperServer) completeWithdrawCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_WITHDRAW,
			ExpectedMethod:    stark.MethodType_T,
			ExpectedStatus:    stark.Status_SUBMITTED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_SUBMIT_FAILED,
		})
}
