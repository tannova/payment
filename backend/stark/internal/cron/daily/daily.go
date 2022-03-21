package daily

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

type Daily interface{}

type daily struct {
	entClient   *ent.Client
	cron        *cron.Cron
	cronEntryID cron.EntryID
	logger      *zap.Logger
}

func New(dbEnt *ent.Client) Daily {
	logger := logging.Logger(nil)
	d := &daily{
		entClient: dbEnt,
		cron:      cron.New(),
		logger:    logger,
	}

	var err error
	d.cronEntryID, err = d.cron.AddFunc("@daily", d.run)
	if err != nil {
		logger.Error("add cron job func error", zap.Error(err))
	} else {
		d.cron.Start()
	}

	return d
}

func (d daily) run() {
	ctx := context.Background()
	_, err := d.entClient.SystemBankAccount.Update().
		SetDailyUsedAmount(0).
		SetDailyBalance(0).
		SetLastUpdatedBalance(time.Now().UTC()).
		Save(ctx)
	if err != nil {
		d.logger.Error("failed to reset banks balance and used amount", zap.Error(err))
	}

	_, err = d.entClient.SystemEWallet.Update().
		SetDailyUsedAmount(0).
		SetDailyBalance(0).
		SetLastUpdatedBalance(time.Now().UTC()).
		Save(ctx)
	if err != nil {
		d.logger.Error("failed to reset e wallets balance and used amount", zap.Error(err))
	}
}
