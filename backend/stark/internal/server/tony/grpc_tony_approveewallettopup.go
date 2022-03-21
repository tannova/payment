package tony

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *tonyServer) ApproveEWalletTopUp(
	ctx context.Context,
	request *stark.ApproveEWalletTopUpRequest,
) (*stark.ApproveEWalletTopUpReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	metaData, err := utils.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}
	outCtx := metadata.NewOutgoingContext(ctx, metaData)

	expectedType := int32(stark.PaymentType_TOPUP)
	expectedMethod := int32(stark.MethodType_E)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)
	nextFailedStatus := int32(stark.Status_APPROVE_FAILED)
	now := time.Now().UTC()

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		payment, err := tx.Client().Payment.
			Query().
			Where(
				ep.IDEQ(request.GetPaymentId()),
				ep.TypeEQ(expectedType),
				ep.MethodEQ(expectedMethod),
				ep.StatusEQ(expectedStatus),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("get payment error", zap.Error(err))
			return err
		}

		paymentDetail, err := payment.QueryPaymentEWalletDetail().Only(ctx)
		if err != nil {
			logger.Error("get payment detail error", zap.Error(err))
			return err
		}

		makePaymentReply, err := s.blackWidowClient.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     paymentDetail.Amount,
			Type:       natasha.PaymentType_USER_TOP_UP,
		})

		var (
			status int32
			note   string
		)
		if err != nil {
			logger.Error("make payment error", zap.Error(err))
			status = nextFailedStatus
			note = err.Error()
		} else {
			status = nextSuccessStatus
			note = request.GetNote()
		}
		payment, err = tx.Client().Payment.
			UpdateOne(payment).
			SetApprovedBy(tellerID).
			SetApprovedAt(time.Now().UTC()).
			SetUpdatedBy(tellerID).
			SetUpdatedAt(now).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("could not update status payment", zap.Error(err))
			return err
		}

		paymentDetail, err = s.entClient.PaymentEWalletDetail.
			UpdateOne(paymentDetail).
			SetImageURL(request.ImageUrl).
			SetTxID(request.TxId).
			Save(ctx)
		if err != nil {
			logger.Error("update payment detail err", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(now).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetNote(note).
			SetPayment(payment).
			Save(ctx)

		if err != nil {
			logger.Error("create revision err", zap.Error(err))
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
			s.completeApproveEWalletTopupCallback)

		if err != nil {
			logger.Error("complete payment err", zap.Error(err))
			return err
		}

		systemEWallet, err := s.entClient.SystemEWallet.Query().
			Where(esw.And(
				esw.AccountName(paymentDetail.SystemAccountName),
				esw.AccountPhoneNumber(paymentDetail.SystemAccountPhoneNumber),
				esw.EWalletName(paymentDetail.EWalletName),
			)).
			First(ctx)
		if err != nil {
			logger.Error("failed to get system bank", zap.Error(err))
			return err
		}

		_, err = systemEWallet.Update().
			AddDailyUsedAmount(1).
			AddDailyBalance(paymentDetail.Amount).
			SetLastUpdatedBalance(time.Now().UTC()).
			Save(ctx)

		if err != nil {
			logger.Error("failed to update daily info of system account", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.ApproveEWalletTopUpReply{}, nil
}

func (s *tonyServer) completeApproveEWalletTopupCallback(
	request *jarvis.CompletePaymentRequest,
	responseErr error,
) {
	s.completeCallback(
		request,
		responseErr,
		PaymentTransition{
			ExpectedType:      stark.PaymentType_TOPUP,
			ExpectedMethod:    stark.MethodType_E,
			ExpectedStatus:    stark.Status_APPROVED,
			NextSuccessStatus: stark.Status_COMPLETED,
			NextFailedStatus:  stark.Status_APPROVE_FAILED,
		})
}
