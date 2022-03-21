package romanoff

import (
	"context"
	"errors"
	"time"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/repository/tx"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"go.uber.org/zap"
)

func (s *romanoffServer) CreateVoucher(ctx context.Context, request *natasha.CreateVoucherRequest) (*natasha.CreateVoucherReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	voucherType := request.GetType()
	// the amount already make sure > 0 with the validate proto
	amount := int64(request.GetAmount())
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		s.logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	merchantBalances, err := s.repository.Payment.GetMerchantBalances(ctx, request.GetMerchantId(), time.Unix(0, 0).UTC(), time.Unix(0, 0).UTC())
	s.logger.Debug("balances", zap.Any("balances", merchantBalances))
	if err != nil {
		s.logger.Error("can not get merchant balance", zap.Error(err))
		return nil, err
	}
	if voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT ||
		voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS ||
		voucherType == natasha.PaymentType_USER_WITHDRAW {

		// check enough balance for withdraw
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

	var paymentID int64 = -1
	err = s.repository.WithTransaction(ctx, func(ctx context.Context, tx tx.Tx) error {
		// We'll change the balance of the MEX immediately when create two kind of WITHDRAW
		// and then refund if it's cancelled
		if voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT ||
			voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS {
			// create payment
			payment, err := tx.Client().Payment.
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
				return err
			}
			paymentID = payment.ID
		}

		createOne := tx.Client().Voucher.
			Create().
			SetMerchantID(request.GetMerchantId()).
			SetType(int32(voucherType)).
			SetStatus(int32(natasha.VoucherStatus_PROCESSING)).
			SetAmount(amount).
			SetCreatorNote(request.GetNote()).
			SetCreatedAt(time.Now().UTC()).
			SetUpdatedAt(time.Now().UTC()).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID)

		if paymentID > 0 {
			createOne.SetPaymentID(paymentID)
		}
		_, err := createOne.Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.logger.Error("can handle payment and voucher", zap.Error(err))
		return nil, err
	}

	return &natasha.CreateVoucherReply{}, nil
}
