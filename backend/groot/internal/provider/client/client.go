package client

import (
	"context"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
)

type Client interface {
	GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error)
	ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) error
}
