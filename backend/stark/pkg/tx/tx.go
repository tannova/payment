package tx

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func WithTransaction(ctx context.Context, logger *zap.Logger, client *ent.Client, fn Fn) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			if err = tx.Rollback(); err != nil {
				logger.Error("can not rollback transaction", zap.Error(err))
			}
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			if err = tx.Rollback(); err != nil {
				logger.Error("can not rollback transaction", zap.Error(err))
			}
		} else {
			// all good, commit
			if err = tx.Commit(); err != nil {
				logger.Error("can not commit transaction", zap.Error(err))
			}
		}
	}()
	err = fn(ctx, tx)

	return err
}

// Tx is an interface that models the standard transaction in
// `ent/tx`.
//
// To ensure `TxFn` funcs cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
type Tx interface {
	Client() *ent.Client
	OnRollback(f ent.RollbackHook)
	OnCommit(f ent.CommitHook)
}

type Fn func(context.Context, Tx) error
