package romanoff

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/repository/tx"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	evoucher "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"
)

func (s *romanoffServer) CancelVoucher(ctx context.Context, request *natasha.CancelVoucherRequest) (*natasha.CancelVoucherReply, error) {
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
	oldVoucher, err := s.eclient.Voucher.Query().Where(evoucher.ID(request.GetId())).First(ctx)
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
		voucher, err := tx.Client().Voucher.Get(ctx, request.GetId())
		if err != nil {
			return err
		}
		voucherType := natasha.PaymentType(voucher.Type)

		updateOne := tx.Client().Voucher.UpdateOneID(request.GetId()).
			SetUpdatedBy(tellerID).
			SetStatus(int32(natasha.VoucherStatus_CANCELED)).
			SetHandlerNote(request.GetNote()).
			SetHandledBy(tellerID).
			SetUpdatedAt(time.Now().UTC()).
			SetUpdatedBy(tellerID)

		_, err = updateOne.Save(ctx)
		if err != nil {
			return err
		}
		// We need to add balance back when we cancel a voucher
		// TODO(tiennv147) find better option than delete the payment
		if voucher.PaymentID > 0 && (voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT ||
			voucherType == natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS) {
			err := tx.Client().Payment.DeleteOneID(voucher.PaymentID).Exec(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &natasha.CancelVoucherReply{}, nil
}
