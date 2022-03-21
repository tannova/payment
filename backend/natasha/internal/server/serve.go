package server

import (
	"context"

	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	postman "gitlab.com/greyhole/postman/pkg/client"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/server/blackwidow"
	"gitlab.com/mcuc/monorepo/backend/natasha/internal/server/natasha"
	"gitlab.com/mcuc/monorepo/backend/natasha/internal/server/romanoff"
	"gitlab.com/mcuc/monorepo/backend/natasha/internal/service/notify"
	stoken "gitlab.com/mcuc/monorepo/backend/natasha/internal/service/token"
	pb "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/natasha/pkg/config"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

// Serve ...
func Serve(cfg *conf.Config, service nightkit.Service) {
	var (
		postmanClient *postman.PostmanClient
	)

	logger := service.Logger()

	eclient, err := ent.Open("mysql", cfg.Database.Url)
	if err != nil {
		logger.Fatal("failed to connect to mysql", zap.Error(err))
	}
	defer func() {
		if err := eclient.Close(); err != nil {
			panic(err)
		}
	}()
	if err := eclient.Schema.Create(context.Background()); err != nil {
		service.Logger().Fatal("failed creating schema resources", zap.Error(err))
	}

	if cfg.FeatureFlag.EnableHawkeye {
		postmanClient, err = postman.NewPostmanClient(cfg.Hawkeye)
		if err != nil {
			logger.Fatal("can not create postman client", zap.Error(err))
		}
	}

	server := service.Server()
	token, err := stoken.New(logger, cfg.MexTwtTokenSigning)
	if err != nil {
		logger.Fatal("can not create servie token", zap.Error(err))
	}
	notify := notify.NewNotify(logger)

	pb.RegisterNatashaServer(server, natasha.NewServer(service, eclient, postmanClient, token))
	pb.RegisterBlackWidowServer(server, blackwidow.NewServer(service, eclient, notify, postmanClient, cfg))
	pb.RegisterRomanoffServer(server, romanoff.NewServer(service, eclient, postmanClient, token, notify))

	service.Serve()
}
