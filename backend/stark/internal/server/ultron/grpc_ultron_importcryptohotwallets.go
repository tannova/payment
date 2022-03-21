package ultron

import (
	"context"
	"time"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func (s *ultronServer) ImportCryptoHotWallets(
	ctx context.Context,
	request *stark.ImportCryptoHotWalletsRequest,
) (*stark.ImportCryptoHotWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	var walletCreates []*ent.SystemCryptoHotWalletCreate
	for _, w := range request.GetRecords() {
		walletCreate := s.entClient.SystemCryptoHotWallet.
			Create().
			SetStatus(int32(stark.CryptoHotWalletStatus_CRYPTO_HOT_WALLET_STATUS_ACTIVE)).
			SetAddress(w.Address).
			SetCryptoType(int32(w.CryptoType)).
			SetCryptoNetworkType(int32(w.CryptoNetworkType)).
			SetTotalBalance(w.TotalBalance).
			SetBalance(w.Balance).
			SetMerchantID(w.MerchantId).
			SetCreatedAt(time.Now().UTC()).
			SetCreatedBy(tellerID).
			SetUpdatedBy(tellerID)
		walletCreates = append(walletCreates, walletCreate)
	}

	err = s.entClient.SystemCryptoHotWallet.CreateBulk(walletCreates...).Exec(ctx)
	if err != nil {
		logger.Error("can not create bulk wallet", zap.Error(err))
		return nil, err
	}

	return &stark.ImportCryptoHotWalletsReply{
		Records: request.Records,
	}, nil
}
