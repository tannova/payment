package tx

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

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

type TxFn func(context.Context, Tx) error
