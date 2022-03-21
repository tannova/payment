package romanoff

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/repository/tx"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	evoucher "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"
)

func (s *romanoffServer) SubmitVoucher(ctx context.Context, request *natasha.SubmitVoucherRequest) (*natasha.SubmitVoucherReply, error) {
	if err := request.Validate(); err != nil {
		s.logger.Warn("invalid request", zap.Error(err))
		// TODO(nghieppp): define error code
		return nil, err
	}

	_, tellerID, err := getUserID(ctx)
	if err != nil {
		s.logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}
	oldVoucher, err := s.eclient.Voucher.Query().Where(evoucher.IDEQ(request.GetId())).Only(ctx)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("Error DB CancelPayment : %s", err.Error()))
		return nil, err
	}
	if oldVoucher.Status != int32(natasha.VoucherStatus_PROCESSING) {
		err := "status invalid, must be processing"
		s.logger.Error(err)
		return nil, errors.New(err)
	}

	err = s.repository.WithTransaction(ctx, func(ctx context.Context, tx tx.Tx) error {
		// We already created payment for two kind of WITHDRAW, so we don't need to do that again
		voucher, err := tx.Client().Voucher.Get(ctx, request.GetId())
		if err != nil {
			return err
		}
		voucherType := natasha.PaymentType(voucher.Type)
		// Not sure if we need paymentID always link
		var paymentID int64 = -1

		if voucherType != natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT &&
			voucherType != natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS {
			paymentCreate, err := tx.Client().Payment.
				Create().
				SetMerchantID(oldVoucher.MerchantID).
				SetType(int32(oldVoucher.Type)).
				SetAmount(oldVoucher.Amount).
				SetCreatedAt(time.Now().UTC()).
				SetUpdatedAt(time.Now().UTC()).
				SetCreatedBy(tellerID).
				SetUpdatedBy(tellerID).
				Save(ctx)
			if err != nil {
				return err
			}
			paymentID = paymentCreate.ID
		}

		// create voucher
		updateOne := tx.Client().Voucher.UpdateOneID(request.GetId())
		updateOne.SetUpdatedBy(tellerID).
			SetStatus(int32(natasha.VoucherStatus_COMPLETED)).
			SetPayeeProvider(int32(request.GetPayeeProvider())).
			SetPayeeAccount(request.GetPayeeAccount()).
			SetPayeeName(request.GetPayeeName()).
			SetPayerProvider(int32(request.GetPayerProvider())).
			SetPayerAccount(request.GetPayerAccount()).
			SetPayerName(request.GetPayerName()).
			SetTxID(request.GetTxId()).
			SetHandlerNote(request.GetHandlerNote()).
			SetImageURL(request.GetImageUrl()).
			SetHandledBy(tellerID).
			SetUpdatedAt(time.Now().UTC()).
			SetUpdatedBy(tellerID)

		if paymentID > 0 {
			updateOne.SetPaymentID(paymentID)
		}

		_, err = updateOne.Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		// TODO: convert code if needed
		return nil, err
	}

	amount := int64(oldVoucher.Amount)
	merchantBalances, err := s.repository.Payment.GetMerchantBalances(
		ctx,
		oldVoucher.MerchantID,
		time.Unix(0, 0).UTC(),
		time.Unix(0, 0).UTC(),
	)
	if err != nil {
		return nil, fmt.Errorf("can not get mex balance: %w", err)
	}
	curBalance := int64(merchantBalances.Balance) + int64(amount)
	// send Slack once balance reached a certain threshold
	merchant, err := s.repository.Merchant.GetMerchant(ctx, oldVoucher.MerchantID)
	if err == nil && merchant != nil {
		if merchant.ID > 0 {
			msg := &slack.WebhookMessage{
				Text: fmt.Sprintf(
					"The balane is running out. Please add more deposit.\nMerchant id: %v\nCurrent balance: %v",
					oldVoucher.MerchantID,
					merchantBalances.Balance,
				),
			}
			if curBalance <= int64(merchant.SafetyLimit) {
				s.notify.Send(msg, merchant.SlackWebhookURL)
			}
		}
	}

	return &natasha.SubmitVoucherReply{}, nil
}
