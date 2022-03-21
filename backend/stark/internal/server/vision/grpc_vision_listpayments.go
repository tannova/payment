package vision

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ep "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	ebp "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	ecp "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentcryptodetail"
	ewp "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentewalletdetail"
	etp "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

func (s *visionServer) ListPayments(ctx context.Context, request *stark.ListPaymentsRequest) (*stark.ListPaymentsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)

	var (
		limit  = 10
		offset = 0
	)

	if request.GetSize() > 0 {
		limit = int(request.GetSize())
	}
	if request.GetPage() > 0 {
		offset = (int(request.GetPage()) - 1) * limit
	}

	var filters []predicate.Payment

	if request.GetFrom().IsValid() {
		filters = append(filters, ep.UpdatedAtGTE(request.GetFrom().AsTime()))
	}
	if request.GetTo().IsValid() {
		filters = append(filters, ep.UpdatedAtLTE(request.GetTo().AsTime()))
	}

	if len(request.GetPaymentIds()) > 0 {
		ids := make([]int64, len(request.GetPaymentIds()))
		for i, v := range request.GetPaymentIds() {
			ids[i] = v
		}
		filters = append(filters, ep.IDIn(ids...))
	}

	if len(request.GetPaymentTypes()) > 0 {
		paymentTypes := make([]int32, len(request.GetPaymentTypes()))
		for i, v := range request.GetPaymentTypes() {
			paymentTypes[i] = int32(v)
		}
		filters = append(filters, ep.TypeIn(paymentTypes...))
	}

	if len(request.GetMethods()) > 0 {
		methods := make([]int32, len(request.GetMethods()))
		for i, v := range request.GetMethods() {
			methods[i] = int32(v)
		}
		filters = append(filters, ep.TypeIn(methods...))
	}

	if len(request.GetStatuses()) > 0 {
		statuses := make([]int32, len(request.GetStatuses()))
		for i, v := range request.GetStatuses() {
			statuses[i] = int32(v)
		}
		filters = append(filters, ep.StatusIn(statuses...))
	}

	if len(request.GetMerchantIds()) > 0 {
		merchantIDs := make([]int64, len(request.GetMerchantIds()))
		for i, v := range request.GetMerchantIds() {
			merchantIDs[i] = v
		}
		filters = append(filters, ep.MerchantIDIn(merchantIDs...))
	}

	var paymentMethodFilters []predicate.Payment

	if len(request.GetBankNames()) > 0 {
		bankNames := make([]int32, len(request.GetBankNames()))
		for i, v := range request.GetBankNames() {
			bankNames[i] = int32(v)
		}
		paymentMethodFilters = append(paymentMethodFilters, ep.HasPaymentBankingDetailWith(ebp.MerchantUserBankNameIn(bankNames...)))
	}

	if len(request.GetEWalletNames()) > 0 {
		eWalletNames := make([]int32, len(request.GetEWalletNames()))
		for i, v := range request.GetEWalletNames() {
			eWalletNames[i] = int32(v)
		}
		paymentMethodFilters = append(paymentMethodFilters, ep.HasPaymentEWalletDetailWith(ewp.EWalletNameIn(eWalletNames...)))
	}

	if len(request.GetTelcoNames()) > 0 {
		telcoNames := make([]int32, len(request.GetTelcoNames()))
		for i, v := range request.GetTelcoNames() {
			telcoNames[i] = int32(v)
		}
		paymentMethodFilters = append(paymentMethodFilters, ep.HasPaymentTelcoDetailWith(etp.TelcoNameIn(telcoNames...)))
	}

	if len(request.GetCryptoWalletName()) > 0 {
		cryptoWalletName := make([]int32, len(request.GetCryptoWalletName()))
		for i, v := range request.GetCryptoWalletName() {
			cryptoWalletName[i] = int32(v)
		}
		paymentMethodFilters = append(paymentMethodFilters, ep.HasPaymentCryptoDetailWith(ecp.CryptoWalletNameIn(cryptoWalletName...)))
	}

	if len(paymentMethodFilters) > 0 {
		filters = append(filters, ep.Or(
			paymentMethodFilters...,
		))
	}

	query := s.entClient.Payment.
		Query().
		Where(ep.And(filters...)).
		Limit(limit).
		Offset(offset)
	if request.OrderAscUpdatedAt {
		query.Order(ent.Asc(ep.FieldUpdatedAt))
	} else {
		query.Order(ent.Desc(ep.FieldUpdatedAt))
	}

	payments, err := query.All(ctx)
	if err != nil {
		logger.Error("fail to get list payment", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	total, err := query.Count(ctx)
	if err != nil {
		logger.Error("fail to total of list payment", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	records := make([]*stark.PaymentWithDetail, len(payments))
	for i, payment := range payments {
		records[i] = &stark.PaymentWithDetail{
			Payment:       transformer.GetPayment(payment),
			PaymentDetail: transformer.GetPaymentDetail(ctx, payment),
		}
	}

	return &stark.ListPaymentsReply{
		Records:     records,
		Total:       uint64(total),
		CurrentPage: uint32(offset/limit + 1),
	}, nil
}
