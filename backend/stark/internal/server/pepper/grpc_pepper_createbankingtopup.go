package pepper

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) CreateBankingTopUp(ctx context.Context, request *stark.CreateBankingTopUpRequest) (*stark.CreateBankingTopUpReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	_, tellerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	paymentInfo, err := s.getPaymentInfoByPaymentCode.Get(ctx, &stark.GetPaymentInfoByPaymentCodeRequest{
		Code: request.PaymentCode,
	})
	if err != nil {
		logger.Error("failed to get payment info in banking top up", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		status := int32(stark.Status_CREATED)
		method := int32(stark.MethodType_T)
		paymentType := int32(stark.PaymentType_TOPUP)

		payment, err := tx.Client().Payment.Create().
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetMerchantID(paymentInfo.MerchantId).
			SetMerchantUserID(paymentInfo.GetMerchantUserId()).
			SetMethod(method).
			SetType(paymentType).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking top up payment", zap.Error(err))
			return err
		}

		_, err = tx.Client().PaymentBankingDetail.Create().
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetMerchantUserBankName(int32(paymentInfo.GetBankName())).
			SetMerchantUserAccountName(request.MerchantUserAccountName).
			SetMerchantUserAccountNumber(request.MerchantUserAccountNumber).
			SetSystemAccountName(request.SystemAccountName).
			SetSystemAccountNumber(request.SystemAccountNumber).
			SetSystemAccountBankName(int32(paymentInfo.GetBankName())).
			SetAmount(request.Amount).
			SetPaymentCode(request.PaymentCode).
			SetPayment(payment).
			SetMerchantUserID(paymentInfo.MerchantUserId).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking top up payment detail", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetNote(request.GetNote()).
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking top up payment revision", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Internal(err.Error())
	}
	return &stark.CreateBankingTopUpReply{}, nil
}
