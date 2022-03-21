package transaction

import (
	"context"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
)

type Transaction interface {
	Create(ctx context.Context, tx *ent.Transaction) (*ent.Transaction, error)
}

type transaction struct {
	logger    *zap.Logger
	entClient *ent.Client
}

func New(logger *zap.Logger, entClient *ent.Client) Transaction {
	return &transaction{
		logger:    logger,
		entClient: entClient,
	}
}

func (t transaction) Create(ctx context.Context, tx *ent.Transaction) (*ent.Transaction, error) {
	return t.entClient.Transaction.Create().
		SetAmount(tx.Amount).
		SetRequest(tx.Request).
		SetResponse(tx.Response).
		SetProvider(tx.Provider).
		SetTxID(tx.TxID).
		SetStatus(tx.Status).
		SetCode(tx.Code).
		SetSerial(tx.Serial).
		SetCallbackResponse(tx.CallbackResponse).
		Save(ctx)
}
