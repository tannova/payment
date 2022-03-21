package romanoff

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/predicate"
	evoucher "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"
)

const (
	_defaultLimit  = 10
	_defaultOffset = 0
)

func (s *romanoffServer) ListVouchers(ctx context.Context, request *natasha.ListVouchersRequest) (*natasha.ListVouchersReply, error) {
	var (
		limit  = _defaultLimit
		offset = _defaultOffset
	)
	if request.GetSize() > 0 {
		limit = int(request.GetSize())
	}
	if request.GetPage() > 0 {
		offset = (int(request.GetPage()) - 1) * limit
	}

	var paymentTypes []int32
	if len(request.Types) <= 0 {
		paymentTypes = append(paymentTypes,
			int32(natasha.PaymentType_USER_TOP_UP),
			int32(natasha.PaymentType_USER_WITHDRAW),
			int32(natasha.PaymentType_MERCHANT_DEPOSIT_ADDITIONAL),
			int32(natasha.PaymentType_MERCHANT_DEPOSIT_COMPENSATION),
			int32(natasha.PaymentType_MERCHANT_WITHDRAW_PROFIT),
			int32(natasha.PaymentType_MERCHANT_WITHDRAW_FUNDS),
		)
	} else {
		paymentTypes = append(paymentTypes, request.Types...)
	}

	var statuses []int32
	if len(request.Statuses) <= 0 {
		statuses = append(statuses,
			int32(natasha.VoucherStatus_PROCESSING),
			int32(natasha.VoucherStatus_COMPLETED),
			int32(natasha.VoucherStatus_CANCELED),
		)
	} else {
		statuses = append(statuses, request.Statuses...)
	}

	var filters []predicate.Voucher
	filters = append(filters,
		evoucher.MerchantIDEQ(request.GetMerchantId()),
		evoucher.TypeIn(paymentTypes...),
		evoucher.StatusIn(statuses...),
	)
	id := request.GetId()
	if id > 0 {
		filters = append(filters, evoucher.IDEQ(int64(id)))
	}

	vouchers, err := s.eclient.Voucher.
		Query().
		Where(
			evoucher.And(
				filters...,
			),
		).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(evoucher.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}
	records := []*natasha.Voucher{}
	for _, voucher := range vouchers {
		records = append(records, transformer.GetVoucher(voucher))
	}

	return &natasha.ListVouchersReply{
		Records:     records,
		Total:       uint64(len(vouchers)),
		CurrentPage: uint32(offset/limit + 1),
		CurrentSize: uint32(limit),
	}, nil
}
