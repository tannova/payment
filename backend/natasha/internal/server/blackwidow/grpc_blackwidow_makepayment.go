package blackwidow

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

func (s *blackWidowServer) MakePayment(ctx context.Context, request *natasha.MakePaymentRequest) (*natasha.MakePaymentReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	voucherType := request.GetType()
	amount := int64(request.GetAmount())
	if amount <= 0 {
		err := fmt.Sprintf("value must be greater than 0")
		s.logger.Error(err)
		return nil, errors.New(err)
	}
	merchantBalances, err := s.repository.Payment.GetMerchantBalances(ctx, request.GetMerchantId(), time.Unix(0, 0).UTC(), time.Unix(0, 0).UTC())
	currBalance := int64(merchantBalances.Balance)
	s.logger.Debug("balances", zap.Any("balances", merchantBalances))
	if err != nil {
		s.logger.Error("can not get merchant balance", zap.Error(err))
		return nil, err
	}
	if voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT ||
		voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS ||
		voucherType == natasha.PaymentType_USER_WITHDRAW {

		if voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT {
			if merchantBalances.BalanceForMexWithdrawProfit <= 0 || amount > merchantBalances.BalanceForMexWithdrawProfit {
				s.logger.Error("insufficient profit")
				return nil, errors.New("insufficient profit")
			}
		}
		if voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS {
			if merchantBalances.BalanceForMexWithdrawFunds <= 0 || amount > merchantBalances.BalanceForMexWithdrawFunds {
				s.logger.Error("insufficient funds")
				return nil, errors.New("insufficient funds")
			}
		}
		if voucherType == natasha.PaymentType_USER_WITHDRAW {
			if merchantBalances.Balance <= 0 || uint64(amount) > merchantBalances.Balance {
				s.logger.Error("insufficient balance")
				return nil, errors.New("insufficient balance")
			}
		}

		amount = -amount
	}

	_, tellerID, err := getUserID(ctx)
	if err != nil {
		s.logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	payment, err := s.eclient.Payment.
		Create().
		SetMerchantID(request.GetMerchantId()).
		SetType(int32(voucherType)).
		SetAmount(amount).
		SetCreatedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		SetCreatedBy(tellerID).
		SetUpdatedBy(tellerID).
		Save(ctx)
	if err != nil {
		// TODO:(nghieppp) define error code
		return nil, err
	}

	currBalance += int64(amount)

	// send Slack once balance reached a certain threshold
	merchant, err := s.repository.Merchant.GetMerchant(ctx, request.GetMerchantId())
	if err == nil && merchant != nil {
		if merchant.ID > 0 {
			msg := &slack.WebhookMessage{
				Text: fmt.Sprintf(
					"The balane is running out. Please add more deposit.\nMerchant id: %v\nCurrent balance: %v",
					request.GetMerchantId(),
					merchantBalances.Balance,
				),
			}
			if currBalance <= int64(merchant.SafetyLimit) {
				s.notify.Send(msg, merchant.SlackWebhookURL)
			}
		}
	}

	return &natasha.MakePaymentReply{
		Balance: uint64(currBalance),
		Payment: transformer.GetPayment(payment),
	}, nil
}
