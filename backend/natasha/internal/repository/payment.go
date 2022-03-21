package repository

import (
	"context"
	"time"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/service/merchant"
	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/predicate"
)

type Payment interface {
	GetMerchantBalances(ctx context.Context, mexID int64, fromDate, toDate time.Time) (result *merchant.MerchantBalances, err error)
}

type paymentImpl struct {
	dbEnt  *ent.Client
	logger *zap.Logger
}

func NewPayment(logger *zap.Logger, dbEnt *ent.Client) Payment {
	r := &paymentImpl{}
	r.logger = logger
	r.dbEnt = dbEnt
	return r
}

func (impl paymentImpl) GetMerchantBalances(ctx context.Context, mexID int64, fromDate, toDate time.Time) (result *merchant.MerchantBalances, err error) {
	var filters []predicate.Payment
	filters = append(filters, epayment.MerchantIDEQ(mexID))
	if !fromDate.IsZero() && fromDate != time.Unix(0, 0).UTC() {
		filters = append(filters, epayment.CreatedAtGTE(fromDate))
	}
	if !toDate.IsZero() && toDate != time.Unix(0, 0).UTC() {
		filters = append(filters, epayment.CreatedAtLTE(toDate))
	}
	payments, err := impl.dbEnt.Payment.
		Query().
		Where(
			filters...,
		).All(ctx)
	if err != nil {
		impl.logger.Error("Error DB GetMerchantBalances", zap.Error(err))
		return nil, err
	}

	records := []*natasha.Payment{}
	for _, payment := range payments {
		records = append(records, transformer.GetPayment(payment))
	}

	return merchant.GetSumBalances(records), nil
}
