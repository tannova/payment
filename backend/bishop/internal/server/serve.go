package server

import (
	"os"

	"gitlab.com/greyhole/night-kit/pkg/config"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/bishop/internal/server/bishop"
	conf "gitlab.com/mcuc/monorepo/backend/bishop/pkg/config"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/client"
)

// Serve ...
func Run(f *config.Flags) {
	cfg := loadConfig(f)

	logger, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		logging.NewTmpLogger().Fatal("could not instantiate logger", zap.Error(err))
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	morganClient, err := stark.NewMorganClient(cfg.Morgan)
	if err != nil {
		logger.Fatal("failed to new pepper client", zap.Error(err))
	}

	httpServer := bishop.NewHTTPServer(logger, cfg, morganClient)
	logger.Info("start bishop http")
	if err := httpServer.Serve(); err != nil {
		logger.Fatal("failed to serve of http", zap.Error(err))
	}
}

func loadConfig(f *config.Flags) *conf.Config {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logging.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg conf.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	if err := cfg.Validate(); err != nil {
		tmpLogger.Fatal("validating configuration failed", zap.Error(err))
	}

	if f.Validate {
		tmpLogger.Info("configuration validation was successful")
		os.Exit(0)
	}

	return &cfg
}
