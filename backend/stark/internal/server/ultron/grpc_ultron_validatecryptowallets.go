package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	ecw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/cryptowallet"
)

func (s *ultronServer) ValidateCryptoWallets(
	ctx context.Context,
	request *stark.ValidateCryptoWalletsRequest,
) (*stark.ValidateCryptoWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, _, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	var validRecords []*stark.CryptoWallet
	var duplicatedRecords []*stark.CryptoWallet
	for _, w := range request.GetRecords() {
		exists, err := s.entClient.CryptoWallet.Query().
			Where(
				ecw.AddressEQ(w.Address),
				ecw.CryptoTypeEQ(int32(w.CryptoType)),
				ecw.CryptoNetworkTypeEQ(int32(w.CryptoNetworkType)),
			).Exist(ctx)
		if err != nil {
			logger.Error("can not query wallet for validate", zap.Error(err))
			return nil, err
		}
		if exists {
			duplicatedRecords = append(duplicatedRecords, w)
		} else {
			validRecords = append(validRecords, w)
		}
	}

	return &stark.ValidateCryptoWalletsReply{
		ValidRecords:      validRecords,
		DuplicatedRecords: duplicatedRecords,
	}, nil
}
