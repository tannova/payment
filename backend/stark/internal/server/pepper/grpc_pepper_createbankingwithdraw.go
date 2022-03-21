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

func (s *pepperServer) CreateBankingWithdraw(ctx context.Context, request *stark.CreateBankingWithdrawRequest) (*stark.CreateBankingWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	mexID, mexIDStr, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("can not get mex id", zap.Error(err))
		return nil, status.Unauthenticated
	}
	var paymentId int64

	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		caller := getMexIDAudit(mexIDStr)
		status := int32(stark.Status_CREATED)
		method := int32(stark.MethodType_T)
		paymentType := int32(stark.PaymentType_WITHDRAW)

		payment, err := tx.Client().Payment.Create().
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetMerchantID(mexID).
			SetMerchantUserID(request.GetMerchantUserId()).
			SetMethod(method).
			SetType(paymentType).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking withdraw payment", zap.Error(err))
			return err
		}

		_, err = tx.Client().PaymentBankingDetail.Create().
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetMerchantUserBankName(int32(request.BankName)).
			SetMerchantUserAccountName(request.MerchantUserAccountName).
			SetMerchantUserAccountNumber(request.MerchantUserAccountNumber).
			SetAmount(request.Amount).
			SetPaymentCode("").
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking withdraw payment detail", zap.Error(err))
			return err
		}

		_, err = tx.Client().Revision.Create().
			SetCreatedBy(caller).
			SetUpdatedBy(caller).
			SetPaymentID(payment.ID).
			SetStatus(status).
			SetNote("Merchant user created a payment").
			SetPayment(payment).
			Save(ctx)
		if err != nil {
			logger.Error("failed to create banking withdraw payment revision", zap.Error(err))
			return err
		}
		paymentId = payment.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.CreateBankingWithdrawReply{
		PaymentId: paymentId,
	}, nil
}
