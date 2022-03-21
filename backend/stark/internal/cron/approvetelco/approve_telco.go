package approvetelco

import (
	"context"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/telcosetting"

	"github.com/robfig/cron/v3"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/jarvis"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/module/telcopayment"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	epayment "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
)

type ApproveTelco interface{}

type approveTelco struct {
	entClient     *ent.Client
	cron          *cron.Cron
	cronEntryID   cron.EntryID
	logger        *zap.Logger
	natashaCli    natasha.NatashaClient
	jarvisClient  jarvis.Client
	blackWidowCli natasha.BlackWidowClient
	grootCli      groot.GrootClient
	telcoPayment  telcopayment.TelcoPayment
	telcoSetting  telcosetting.Setting
}

func New(
	dbEnt *ent.Client,
	natashaCli natasha.NatashaClient,
	blackWidowCli natasha.BlackWidowClient,
	grootCli groot.GrootClient,
) ApproveTelco {
	logger := logging.Logger(nil)
	jarvisClient := jarvis.New(logger)
	d := &approveTelco{
		entClient:     dbEnt,
		cron:          cron.New(),
		logger:        logger,
		natashaCli:    natashaCli,
		jarvisClient:  jarvisClient,
		blackWidowCli: blackWidowCli,
		grootCli:      grootCli,
		telcoPayment:  telcopayment.New(dbEnt, natashaCli, jarvisClient, blackWidowCli, grootCli),
		telcoSetting:  telcosetting.New(dbEnt),
	}

	var err error
	d.cronEntryID, err = d.cron.AddFunc("@every 10s", d.run)
	if err != nil {
		logger.Error("add cron job func error", zap.Error(err))
	} else {
		d.cron.Start()
	}

	return d
}

func (d approveTelco) run() {
	ctx := context.Background()
	logger := logging.Logger(nil)
	settings, err := d.telcoSetting.GetSettings(ctx, &stark.GetSettingsRequest{})
	if err != nil {
		logger.Error("failed to get seting", zap.Error(err))
		return
	}
	if settings.TopUpAutoApproval == false {
		return
	}

	payments, err := d.entClient.Payment.Query().Where(
		epayment.Status(int32(stark.Status_CREATED)),
		epayment.Type(int32(stark.PaymentType_TOPUP)),
		epayment.Method(int32(stark.MethodType_P)),
	).Limit(10).
		Order(ent.Desc(epayment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		logger.Error("failed to get payments", zap.Error(err))
		return
	}

	for _, payment := range payments {
		_, err := d.telcoPayment.ApproveTopUp(context.Background(), "system", &stark.ApproveTelcoTopUpRequest{
			PaymentId: payment.ID,
			Note:      "Auto Approve payment",
		})
		if err != nil {
			logger.Error("failed to approve payment", zap.Error(err))
			return
		}
	}
}
