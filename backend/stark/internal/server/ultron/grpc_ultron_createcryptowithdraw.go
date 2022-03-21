package ultron

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func (s *ultronServer) CreateCryptoWithdraw(
	ctx context.Context,
	request *stark.CreateCryptoWithdrawRequest,
) (*stark.CreateCryptoWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	mexID, mexIDStr, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	var paymentId int64

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		caller := getMexIDAudit(mexIDStr)
		now := time.Now().UTC()
		status := int32(stark.Status_CREATED)
		method := int32(stark.MethodType_C)
		paymentType := int32(stark.PaymentType_WITHDRAW)

		payment, err := tx.Client().Payment.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetMerchantID(mexID).
			SetMerchantUserID(request.MerchantUserId).
			SetMethod(method).
			SetType(paymentType).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("create payment error", zap.Error(err))
			return err
		}

		_, err = tx.Client().PaymentCryptoDetail.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetPaymentID(payment.ID).
			SetCryptoType(int32(request.CryptoType)).
			SetCryptoNetworkType(int32(request.CryptoNetworkType)).
			SetCryptoWalletName(int32(request.CryptoWalletName)).
			SetReceiverAddress(request.Address).
			SetAmount(request.Amount).
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("create payment detail error", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(now).
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetNote("Merchant create crypto withdraw payment").
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("could create revision", zap.Error(err))
			return err
		}

		paymentId = payment.ID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.CreateCryptoWithdrawReply{
		PaymentId: paymentId,
	}, nil
}
