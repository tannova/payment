package ultron

import (
	"context"
	"time"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *ultronServer) ApproveCryptoWithdraw(
	ctx context.Context,
	request *stark.ApproveCryptoWithdrawRequest,
) (*stark.ApproveCryptoWithdrawReply, error) {
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
	var currentPayment *ent.Payment
	var currentPaymentDetail *ent.PaymentCryptoDetail

	expectedType := int32(stark.PaymentType_WITHDRAW)
	expectedMethod := int32(stark.MethodType_C)
	expectedStatus := int32(stark.Status_CREATED)
	nextSuccessStatus := int32(stark.Status_APPROVED)

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

		paymentDetail, err := payment.QueryPaymentCryptoDetail().Only(ctx)
		if err != nil {
			logger.Error("get payment detail error", zap.Error(err))
			return err
		}

		payment, err = payment.Update().
			SetApprovedBy(tellerID).
			SetApprovedAt(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			logger.Error("failed to get update approve by", zap.Error(err))
			return err
		}
		_, err = s.blackWidowClient.MakePayment(outCtx, &natasha.MakePaymentRequest{
			MerchantId: payment.MerchantID,
			Amount:     convertUSDT2VND(paymentDetail.Amount),
			Type:       natasha.PaymentType_USER_WITHDRAW,
		})

		if err != nil {
			logger.Error("failed to make payment", zap.Error(err))
			return err
		}

		paymentUpdate, err := tx.Client().Payment.
			UpdateOne(payment).
			SetUpdatedBy(tellerID).
			SetUpdatedAt(time.Now().UTC()).
			SetStatus(nextSuccessStatus).
			Save(ctx)
		if err != nil {
			logger.Error("could not update status payment", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetPaymentID(payment.ID).
			SetStatus(nextSuccessStatus).
			SetNote(request.GetNote()).
			SetPayment(payment).
			Save(ctx)

		if err != nil {
			logger.Error("create revision err", zap.Error(err))
			return err
		}

		reply, err := s.natashaClient.GetMerchant(outCtx, &natasha.GetMerchantRequest{
			Id: payment.MerchantID,
		})
		if err != nil {
			logger.Error("get merchant err", zap.Error(err))
			return err
		}
		completeRequest, err := createCompletePaymentRequest(
			ctx,
			paymentUpdate,
			paymentDetail,
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
			// empty callback
			func(*jarvis.CompletePaymentRequest, error) {})

		if err != nil {
			logger.Error("complete payment err", zap.Error(err))
			return err
		}
		currentPayment = payment
		currentPaymentDetail = paymentDetail

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Start goroutine for auto transfer and submit
	if setting, err := s.GetCryptoSettings(ctx, &stark.GetCryptoSettingsRequest{}); err == nil {
		if setting.AutoTransferWithdrawCrypto {
			go s.transferCryptoAndSubmitWithdraw(currentPayment, currentPaymentDetail)
		}
	}

	return &stark.ApproveCryptoWithdrawReply{}, nil
}
