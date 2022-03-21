package thuthecao

import (
	"fmt"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/client"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

const (
	_providerName = "thuthecao"
)

type config struct{
	baseUrl string
	apiKey string
	callbackUrl string
}

type provider struct {
	cfg *config
	repository  *repository.Repository
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

	apiKey, ok := p.Data["apiKey"]
	if !ok {
		return nil, fmt.Errorf("failed to get apiKey")
	}

	callbackUrl, ok := p.Data["callbackUrl"]
	if !ok {
		return nil, fmt.Errorf("failed to get callbackUrl")
	}

	return &config{
		apiKey: apiKey,
		baseUrl: url,
		callbackUrl: callbackUrl,
	}, nil
}
