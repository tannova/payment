package server

import (
	"context"

	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"go.uber.org/zap"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/client"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/client"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/cron/approvetelco"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/cron/daily"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/getpaymentinfobypaymentcode"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/howard"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/morgan"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/pepper"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/tony"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/ultron"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/server/vision"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	pb "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/stark/pkg/config"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

// Serve ...
func Serve(cfg *conf.Config, service nightkit.Service) {
	// Init database ent
	entClient, err := ent.Open("mysql", cfg.Database.Url)
	if err != nil {
		service.Logger().Fatal("failed to connect to mysql", zap.Error(err))
	}
	defer func() {
		if err := entClient.Close(); err != nil {
			panic(err)
		}
	}()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		service.Logger().Fatal("failed creating schema resources", zap.Error(err))
	}

	natashaCli, err := natasha.NewClientNatasha(cfg.Natasha)
	if err != nil {
		service.Logger().Fatal("failed to new natasha client", zap.Error(err))
	}

	blackWidowCli, err := natasha.NewClientBlackWidow(cfg.Blackwidow)
	if err != nil {
		service.Logger().Fatal("failed to new blackwidow client", zap.Error(err))
	}

	grootCli, err := groot.NewGrootClient(cfg.Groot)
	if err != nil {
		service.Logger().Fatal("failed to new groot client", zap.Error(err))
	}

	setting := setting.NewSettingImpl(service.Logger(), cfg)

	daily.New(entClient)

	approvetelco.New(entClient, natashaCli, blackWidowCli, grootCli)

	getPaymentInfoByPaymentCode := getpaymentinfobypaymentcode.New(setting)

	server := service.Server()
	pb.RegisterHowardServer(server, howard.NewServer(service, entClient))
	pb.RegisterMorganServer(server, morgan.NewServer(service, entClient, natashaCli, blackWidowCli, grootCli))
	pb.RegisterPepperServer(server, pepper.NewServer(service, entClient, natashaCli, setting, getPaymentInfoByPaymentCode, blackWidowCli, natashaCli))
	pb.RegisterTonyServer(server, tony.NewServer(service, entClient, natashaCli, blackWidowCli, getPaymentInfoByPaymentCode, setting))
	pb.RegisterUltronServer(server, ultron.NewServer(service, setting, entClient, natashaCli, blackWidowCli))
	pb.RegisterVisionServer(server, vision.NewServer(service, entClient, setting, getPaymentInfoByPaymentCode))
	service.Serve()
}
