package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	pes "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemcryptohotwallet"
)

func (s *ultronServer) GetSystemCryptoHotWallets(
	ctx context.Context,
	request *stark.GetSystemCryptoHotWalletsRequest,
) (*stark.GetSystemCryptoHotWalletsReply, error) {

	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)

	wallets, err := s.entClient.SystemCryptoHotWallet.
		Query().
		Where(pes.And(
			pes.CryptoTypeEQ(int32(request.CryptoType)),
			pes.CryptoNetworkTypeEQ(int32(request.CryptoNetworkType)),
			pes.MerchantIDEQ(request.MerchantId),
			pes.BalanceGTE(request.Amount),
		)).
		Order(ent.Desc(pes.FieldBalance)).
		All(ctx)
	if err != nil {
		logger.Error("could not query crypto wallets", zap.Error(err))
		return nil, err
	}

	records := []*stark.SystemCryptoHotWallet{}
	for _, wallet := range wallets {
		records = append(records, transformer.GetSystemCryptoHotWallet(wallet))
	}

	return &stark.GetSystemCryptoHotWalletsReply{
		Records: records,
	}, nil
}
