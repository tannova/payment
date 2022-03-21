package repository

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

type Merchant interface {
	GetMerchant(ctx context.Context, id int64) (*ent.Merchant, error)
}

type merchantImpl struct {
	dbEnt  *ent.Client
	logger *zap.Logger
}

func NewMerchant(logger *zap.Logger, dbEnt *ent.Client) Merchant {
	r := &merchantImpl{}
	r.logger = logger
	r.dbEnt = dbEnt
	return r
}

func (impl merchantImpl) GetMerchant(ctx context.Context, id int64) (*ent.Merchant, error) {
	result, err := impl.dbEnt.Merchant.Get(ctx, id)
	if err != nil {
		impl.logger.Warn(fmt.Sprintf("Error DB GetMerchant : %s", err.Error()))
		return nil, err
	}

	return result, nil
}

// TODO:(nghieppp) move to repository
