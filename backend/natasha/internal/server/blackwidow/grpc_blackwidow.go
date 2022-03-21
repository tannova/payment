package blackwidow

import (
	"context"
	"strconv"

	"github.com/DataDog/datadog-go/statsd"
	"gitlab.com/greyhole/myid/pkg/extractor"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/greyhole/postman/pkg/client"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/repository"
	snotify "gitlab.com/mcuc/monorepo/backend/natasha/internal/service/notify"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/natasha/pkg/config"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

func NewServer(service nightkit.Service, eclient *ent.Client, notify snotify.Notify, postman *client.PostmanClient, cfg 		*conf.Config) natasha.BlackWidowServer {
	return &blackWidowServer{
		logger:     service.Logger(),
		stats:      service.Stats(),
		eclient:    eclient,
		notify:     notify,
		postman:    postman,
		repository: repository.NewRepository(service.Logger(), eclient),
		cfg: cfg,
	}
}

type blackWidowServer struct {
	natasha.UnimplementedBlackWidowServer

	logger     *zap.Logger
	stats      *statsd.Client
	eclient    *ent.Client
	notify     snotify.Notify
	repository repository.Repository
	postman    *client.PostmanClient
	cfg 		*conf.Config
}

func getUserID(ctx context.Context) (int64, string, error) {
	n, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return 0, "", err
	}
	return n, strconv.FormatInt(n, 10), nil
}
