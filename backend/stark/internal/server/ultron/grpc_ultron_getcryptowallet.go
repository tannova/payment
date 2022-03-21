package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ecw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/cryptowallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/tx"
)

func (s *ultronServer) GetCryptoWallet(
	ctx context.Context,
	request *stark.GetCryptoWalletRequest,
) (*stark.GetCryptoWalletReply, error) {
	const tag = "GetCryptoWallet"
	logger := logging.Logger(ctx)
	if err := request.Validate(); err != nil {
		return nil, err
	}

	mexID, _, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	var address string
	err = tx.WithTransaction(ctx, logger, s.entClient, func(ctx context.Context, tx tx.Tx) error {
		wallet, err := tx.Client().CryptoWallet.
			Query().
			Where(
				ecw.CryptoTypeEQ(int32(request.CryptoType)),
				ecw.CryptoNetworkTypeEQ(int32(request.CryptoNetworkType)),
				ecw.Status(int32(stark.CryptoWalletStatus_AVAILABLE)),
			).
			ForUpdate().
			First(ctx)
		if err != nil {
			logger.Error("could not get wallet", zap.Error(err))
			return err
		}
		uWallet, err := tx.Client().CryptoWallet.UpdateOne(wallet).
			SetMerchantID(mexID).
			SetMerchantUserID(request.MerchantUserId).
			SetStatus(int32(stark.CryptoWalletStatus_USED)).
			Save(ctx)
		if err != nil {
			logger.Error("could not update wallet balances", zap.Error(err))
			return err
		}
		address = uWallet.Address

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stark.GetCryptoWalletReply{
		Address: address,
	}, nil
}
