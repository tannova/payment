package tony

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
)

func (s *tonyServer) GetSystemEWallets(
	ctx context.Context,
	request *stark.GetSystemEWalletsRequest,
) (*stark.GetSystemEWalletsReply, error) {
	wallets, err := s.entClient.SystemEWallet.Query().
		Where(esw.And(
			esw.EWalletName(int32(request.EWalletName)),
			esw.Status(int32(stark.EWalletStatus_EWALLET_ACTIVE)),
			esw.MerchantID(request.MerchantId),
		)).
		Order(ent.Asc(
			esw.FieldDailyUsedAmount,
		)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return &stark.GetSystemEWalletsReply{
		SystemEWallets: transformer.GetSystemEWallets(wallets),
	}, nil
}
