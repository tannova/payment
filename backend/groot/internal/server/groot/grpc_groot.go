package groot

import (
	"github.com/DataDog/datadog-go/statsd"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/client"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

func NewServer(service nightkit.Service, repository *repository.Repository, featureFlag *conf.FeatureFlag, processor client.Client) groot.GrootServer {
	return &grootServer{
		logger:      service.Logger(),
		stats:       service.Stats(),
		repository:  repository,
		featureFlag: featureFlag,
		processor: processor,
	}
}

type grootServer struct {
	logger      *zap.Logger
	stats       *statsd.Client
	repository  *repository.Repository
	featureFlag *conf.FeatureFlag
	processor client.Client
	groot.UnimplementedGrootServer
}
