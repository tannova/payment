package tony

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func (s *tonyServer) CreateEWalletTopUp(ctx context.Context, request *stark.CreateEWalletTopUpRequest) (*stark.CreateEWalletTopUpReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}
	paymentInfo, err := s.getPaymentInfoByPaymentCode.Get(ctx, &stark.GetPaymentInfoByPaymentCodeRequest{
		Code: request.PaymentCode,
	})
	if err != nil {
		logger.Error("get payment info by code error", zap.Error(err))
		return nil, err
	}

	var paymentId int64

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		status := int32(stark.Status_CREATED)
		method := int32(stark.MethodType_E)
		paymentType := int32(stark.PaymentType_TOPUP)
		now := time.Now().UTC()

		payment, err := tx.Client().Payment.Create().
			SetCreatedAt(now).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetMerchantID(paymentInfo.GetMerchantId()).
			SetMerchantUserID(paymentInfo.GetMerchantUserId()).
			SetMethod(method).
			SetType(paymentType).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("create payment error", zap.Error(err))
			return err
		}

		_, err = tx.Client().PaymentEWalletDetail.Create().
			SetCreatedAt(now).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetPaymentCode(request.GetPaymentCode()).
			SetMerchantUserID(paymentInfo.MerchantUserId).
			SetEWalletName(int32(paymentInfo.GetEWalletName())).
			SetMerchantUserAccountPhoneNumber(request.GetMerchantUserAccountPhoneNumber()).
			SetMerchantUserAccountName(request.GetMerchantUserAccountName()).
			SetSystemAccountName(request.SystemAccountName).
			SetSystemAccountPhoneNumber(request.SystemAccountPhoneNumber).
			SetAmount(request.GetAmount()).
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("create payment detail error", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedAt(now).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetNote(request.GetNote()).
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("create payment revision error", zap.Error(err))
			return err
		}

		paymentId = payment.ID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.CreateEWalletTopUpReply{
		PaymentId: paymentId,
	}, nil
}
