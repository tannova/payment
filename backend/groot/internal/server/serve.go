package server

import (
	"context"

	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/processor"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/server/groot"
	pb "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

// Serve ...
func Serve(cfg *conf.Config, service nightkit.Service) {
	logger := service.Logger()

	r, err := repository.Connect(logger, cfg)
	if err != nil {
		logger.Fatal("failed to connect to repositories", zap.Error(err))
	}

	if err = r.Client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema", zap.Error(err))
	}

	processor, err := processor.New(cfg, r)

	defer r.Close()

	server := service.Server()
	pb.RegisterGrootServer(server, groot.NewServer(service, r, cfg.FeatureFlag, processor))
	service.Serve()
}
