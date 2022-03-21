package banthe247

import (
	"context"
	"fmt"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/client"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

const (
	_providerName = "banthe247"
)

type config struct{
	baseURL string
	security string
	userName string
	password string
}

type provider struct {
	cfg *config
	repository  *repository.Repository
}

func (p provider) ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) error {
	panic("implement me")
}

func New(cfg *conf.Config, repository  *repository.Repository) (client.Client, error) {
	p, ok := cfg.Providers[_providerName]
	if !ok {
		return nil, fmt.Errorf("failed to get provider config")
	}

	c, err := loadConfig(p)
	if err != nil {
		return nil, err
	}

	return &provider{
		cfg: c,
		repository: repository,
	}, nil
}

func loadConfig(p *conf.Provider) (*config, error) {
	url, ok := p.Data["url"]
	if !ok {
		return nil, fmt.Errorf("failed to get url")
	}

	userName, ok := p.Data["userName"]
	if !ok {
		return nil, fmt.Errorf("failed to get userName")
	}

	password, ok := p.Data["password"]
	if !ok {
		return nil, fmt.Errorf("failed to get password")
	}

	security, ok := p.Data["security"]
	if !ok {
		return nil, fmt.Errorf("failed to get password")
	}

	return &config{
		baseURL: url,
		userName: userName,
		password: password,
		security: security,
	}, nil
}
