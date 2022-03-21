package repository

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/repository/tx"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

type Repository struct {
	Merchant Merchant
	Payment  Payment
	Client   *ent.Client

	logger *zap.Logger
}

func NewRepository(logger *zap.Logger, client *ent.Client) Repository {
	r := Repository{
		Client: client,
		logger: logger,
	}
	r.Merchant = NewMerchant(logger, client)
	r.Payment = NewPayment(logger, client)
	return r
}

// A TxFn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (r *Repository) WithTransaction(ctx context.Context, fn tx.TxFn) (err error) {
	txx, err := r.Client.Tx(ctx)
	if err != nil {
		r.logger.Error("can not begin transaction", zap.Error(err))
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			if err := txx.Rollback(); err != nil {
				r.logger.Error("can not rollback transaction", zap.Error(err))
			}
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			if err := txx.Rollback(); err != nil {
				r.logger.Error("can not rollback transaction", zap.Error(err))
			}
		} else {
			// all good, commit
			err = txx.Commit()
		}
	}()

	return fn(ctx, txx)
}
