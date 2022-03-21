package telco

import (
	"context"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
	entTelco "gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/telco"
)

type Telco interface {
	GetAvailable(ctx context.Context, branch int64, amount uint64) (*ent.Telco, error)
}

type telco struct {
	logger    *zap.Logger
	entClient *ent.Client
}

func New(logger *zap.Logger, entClient *ent.Client) Telco {
	return &telco{
		logger:    logger,
		entClient: entClient,
	}
}

func (t telco) GetAvailable(ctx context.Context, branch int64, amount uint64) (*ent.Telco, error) {
	card, err := t.entClient.Telco.Query().
		Where(entTelco.And(
			entTelco.TelcoName(branch),
			entTelco.Amount(amount),
			entTelco.Used(false),
		)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	card, err = t.entClient.Telco.UpdateOne(card).SetUsed(true).Save(ctx)
	if err != nil {
		return nil, err
	}

	return card, nil
}
