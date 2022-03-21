package ultron

import (
	"context"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *ultronServer) LoadCryptoWallets(ctx context.Context, request *stark.LoadCryptoWalletsRequest) (*stark.LoadCryptoWalletsReply, error) {
	return &stark.LoadCryptoWalletsReply{
		Data: `{"status": "Success"}`,
	}, nil
}
