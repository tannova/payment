package naptudong

import (
	"fmt"

	"gitlab.com/mcuc/monorepo/backend/groot/internal/provider/client"
	"gitlab.com/mcuc/monorepo/backend/groot/internal/repository"
	conf "gitlab.com/mcuc/monorepo/backend/groot/pkg/config"
)

const (
	_providerName = "naptudong"
)

type config struct{
	baseUrl string
	partnerID string
	partnerKey string
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

	partnerID, ok := p.Data["partnerID"]
	if !ok {
		return nil, fmt.Errorf("failed to get apiKey")
	}

	partnerKey, ok := p.Data["partnerKey"]
	if !ok {
		return nil, fmt.Errorf("failed to get callbackUrl")
	}

	return &config{
		partnerID: partnerID,
		baseUrl: url,
		partnerKey: partnerKey,
	}, nil
}
