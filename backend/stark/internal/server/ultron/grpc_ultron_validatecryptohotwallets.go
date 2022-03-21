package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	pes "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemcryptohotwallet"
)

func (s *ultronServer) ValidateCryptoHotWallets(
	ctx context.Context,
	request *stark.ValidateCryptoHotWalletsRequest,
) (*stark.ValidateCryptoHotWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, _, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	var validRecords []*stark.SystemCryptoHotWallet
	var duplicatedRecords []*stark.SystemCryptoHotWallet
	for _, w := range request.GetRecords() {
		exists, err := s.entClient.SystemCryptoHotWallet.Query().
			Where(
				pes.AddressEQ(w.Address),
				pes.CryptoTypeEQ(int32(w.CryptoType)),
				pes.CryptoNetworkTypeEQ(int32(w.CryptoNetworkType)),
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

	return &stark.ValidateCryptoHotWalletsReply{
		ValidRecords:      validRecords,
		DuplicatedRecords: duplicatedRecords,
	}, nil
}
