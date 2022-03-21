package vision

import (
	"github.com/DataDog/datadog-go/statsd"
	nightkit "gitlab.com/greyhole/night-kit/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/getpaymentinfobypaymentcode"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/setting"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func NewServer(service nightkit.Service,
	entClient *ent.Client, setting setting.Setting,
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode,
) stark.VisionServer {
	return &visionServer{
		logger:                      service.Logger(),
		stats:                       service.Stats(),
		entClient:                   entClient,
		setting:                     setting,
		getPaymentInfoByPaymentCode: getPaymentInfoByPaymentCode,
	}
}

type visionServer struct {
	stark.UnimplementedVisionServer

	logger                      *zap.Logger
	stats                       *statsd.Client
	entClient                   *ent.Client
	setting                     setting.Setting
	getPaymentInfoByPaymentCode getpaymentinfobypaymentcode.GetPaymentInfoByPaymentCode
}
