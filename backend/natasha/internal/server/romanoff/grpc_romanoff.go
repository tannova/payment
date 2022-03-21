package romanoff

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
	stoken "gitlab.com/mcuc/monorepo/backend/natasha/internal/service/token"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent"
)

func NewServer(
	service nightkit.Service,
	eclient *ent.Client,
	postman *client.PostmanClient,
	token stoken.Token,
	notify snotify.Notify,
) natasha.RomanoffServer {
	return &romanoffServer{
		logger:     service.Logger(),
		stats:      service.Stats(),
		eclient:    eclient,
		postman:    postman,
		token:      token,
		notify:     notify,
		repository: repository.NewRepository(service.Logger(), eclient),
	}
}

type romanoffServer struct {
	natasha.UnimplementedRomanoffServer

	logger     *zap.Logger
	stats      *statsd.Client
	eclient    *ent.Client
	extractor  extractor.Extractor
	postman    *client.PostmanClient
	token      stoken.Token
	notify     snotify.Notify
	repository repository.Repository
}

func getUserID(ctx context.Context) (int64, string, error) {
	n, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return 0, "", err
	}
	return n, strconv.FormatInt(n, 10), nil
}
