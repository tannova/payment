package ultron

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func (s *ultronServer) ImportCryptoWallets(ctx context.Context, request *stark.ImportCryptoWalletsRequest) (*stark.ImportCryptoWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	var wallets []*stark.CryptoWallet
	var walletCreates []*ent.CryptoWalletCreate
	for _, w := range request.GetWallets() {
		walletCreate := s.entClient.CryptoWallet.
			Create().
			SetStatus(int32(stark.CryptoWalletStatus_AVAILABLE)).
			SetAddress(w.Address).
			SetCryptoType(int32(w.CryptoType)).
			SetCryptoNetworkType(int32(w.CryptoNetworkType)).
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID)
		walletCreates = append(walletCreates, walletCreate)
	}

	err = s.entClient.CryptoWallet.CreateBulk(walletCreates...).Exec(ctx)
	if err != nil {
		logger.Error("can not create bulk wallet", zap.Error(err))
		return nil, err
	}

	return &stark.ImportCryptoWalletsReply{
		Wallets: wallets,
	}, nil
}
