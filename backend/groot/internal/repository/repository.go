package repository

import (
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository/transaction"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository/telco"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
)

type Repository struct {
	logger *zap.Logger
	Client *ent.Client
	Telco  telco.Telco
	Transaction transaction.Transaction
}

func Connect(logger *zap.Logger, config *conf.Config) (*Repository, error) {
	client, err := ent.Open("mysql", config.Database.Url)
	if err != nil {
		logger.Error("can not connect to mysql", zap.Error(err))
		return nil, err
	}

	return &Repository{
		Client: client,
		Telco:  telco.New(logger, client),
		Transaction: transaction.New(logger, client),
		logger: logger,
	}, nil
}

func (r *Repository) Close() {
	if err := r.Client.Close(); err != nil {
		r.logger.Error("can not close ent connection", zap.Error(err))
	}
}
