package blackwidow

import (
	"context"
	"time"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/predicate"
)

const (
	_defaultLimit  = 10
	_defaultOffset = 0
)

func (s *blackWidowServer) ListPayments(ctx context.Context, request *natasha.ListPaymentsRequest) (*natasha.ListPaymentsReply, error) {
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

	var filters []predicate.Payment
	filters = append(filters, epayment.MerchantIDEQ(request.GetMerchantId()), epayment.TypeIn(paymentTypes...))
	id := request.GetId()
	if id > 0 {
		filters = append(filters, epayment.IDEQ(int64(id)))
	}

	fromDate := request.GetFromDate().AsTime()
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	if !fromDate.IsZero() && fromDate != time.Unix(0, 0).UTC() {
		firstOfMonth = time.Date(fromDate.Year(), fromDate.Month(), 1, 0, 0, 0, 0, fromDate.Location())
	}
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	filters = append(filters, epayment.CreatedAtGTE(firstOfMonth))
	filters = append(filters, epayment.CreatedAtLTE(lastOfMonth))
	// s.logger.Debug("filter firstOfMonth", zap.Any("firstOfMonth", firstOfMonth))
	// s.logger.Debug("filter lastOfMonth", zap.Any("lastOfMonth", lastOfMonth))

	payments, err := s.eclient.Payment.
		Query().
		Where(
			epayment.And(
				filters...,
			),
		).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(epayment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}

	all, err := s.eclient.Merchant.Query().All(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}
	records := []*natasha.Payment{}
	for _, payment := range payments {
		records = append(records, transformer.GetPayment(payment))
	}
	return &natasha.ListPaymentsReply{
		Records:     records,
		Total:       uint64(len(all)),
		CurrentPage: uint32(offset/limit + 1),
		CurrentSize: uint32(limit),
	}, nil
}
