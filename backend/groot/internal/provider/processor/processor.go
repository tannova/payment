package processor

import (
	"context"
	"fmt"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/banthe247"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/client"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/naptudong"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/thenhanh365"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/thuthecao"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

type ChargeCardRequest struct {
	Card *groot.Card
	PaymentID string
	Providers []string
}

func New(cfg *conf.Config, repository *repository.Repository) (client.Client, error) {
	banthe247Client, err := banthe247.New(cfg, repository)
	if err != nil {
		return nil, err
	}

	thuthecaoClient, err := thuthecao.New(cfg, repository)
	if err != nil {
		return nil, err
	}

	thenhanh365Client, err := thenhanh365.New(cfg, repository)
	if err != nil {
		return nil, err
	}

	napTuDongClient, err := naptudong.New(cfg, repository)
	if err != nil {
		return nil, err
	}

	mChargeCardClients := map[string]client.Client {
		"thenhanh365": thenhanh365Client,
		"thuthecao": thuthecaoClient,
		"naptudong": napTuDongClient,
	}

	mGetCardClients := map[string]client.Client {
		"banthe247": banthe247Client,
	}

	return &provider{
		cfg: cfg,
		banthe247: banthe247Client,
		chargeCardClients: mChargeCardClients,
		getCardClients: mGetCardClients,
	}, nil
}

type provider struct {
	cfg *conf.Config
	repository  *repository.Repository
	banthe247 client.Client
	chargeCardClients map[string]client.Client
	getCardClients map[string]client.Client
}

func (p provider) ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) error {
	for _, provider := range request.Providers {
		if c, ok := p.chargeCardClients[provider]; ok {
			logging.Logger(ctx).Info("Provider ", zap.String("provider", provider))
			return c.ChargeCard(ctx, request)
		}
	}
	return nil
}

func (p provider) GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	for _, provider := range request.Providers {
		if c, ok := p.getCardClients[provider]; ok {
			logging.Logger(ctx).Info("Provider ", zap.String("provider", provider))
			return c.GetCard(ctx, request)
		}
	}

	return nil, fmt.Errorf("unsupported providers")
}
