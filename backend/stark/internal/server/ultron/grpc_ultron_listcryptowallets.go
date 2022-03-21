package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ecw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/cryptowallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

func (s *ultronServer) ListCryptoWallets(
	ctx context.Context,
	request *stark.ListCryptoWalletsRequest,
) (*stark.ListCryptoWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)
	_, _, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	var (
		limit  = 10
		offset = 0
	)

	if request.GetSize() > 0 {
		limit = int(request.GetSize())
	}
	if request.GetPage() > 0 {
		offset = (int(request.GetPage()) - 1) * limit
	}

	var filters []predicate.CryptoWallet
	addresses := request.GetAddresses()
	if len(addresses) > 0 {
		filters = append(filters, ecw.AddressIn(addresses...))
	}

	cryptoTypes := request.GetCryptoTypes()
	if len(cryptoTypes) > 0 {
		in := []int32{}
		for _, v := range cryptoTypes {
			in = append(in, int32(v))
		}
		filters = append(filters, ecw.CryptoTypeIn(in...))
	}

	cryptoNetworkTypes := request.GetCryptoNetworkTypes()
	if len(cryptoNetworkTypes) > 0 {
		in := []int32{}
		for _, v := range cryptoNetworkTypes {
			in = append(in, int32(v))
		}
		filters = append(filters, ecw.CryptoNetworkTypeIn(in...))
	}

	statuses := request.GetStatuses()
	if len(statuses) > 0 {
		in := []int32{}
		for _, v := range statuses {
			in = append(in, int32(v))
		}
		filters = append(filters, ecw.StatusIn(in...))
	}

	merchantIDs := request.GetMerchantIds()
	if len(merchantIDs) > 0 {
		filters = append(filters, ecw.MerchantIDIn(merchantIDs...))
	}

	wallets, err := s.entClient.CryptoWallet.
		Query().
		Where(ecw.And(filters...)).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(ecw.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		logger.Error("could not query crypto wallets", zap.Error(err))
		return nil, err
	}
	total, err := s.entClient.CryptoWallet.
		Query().
		Where(ecw.And(filters...)).
		Count(ctx)
	if err != nil {
		logger.Error("could not count crypto wallets", zap.Error(err))
		return nil, err
	}
	records := []*stark.CryptoWallet{}
	for _, wallet := range wallets {
		records = append(records, transformer.GetCryptoWallet(wallet))
	}

	return &stark.ListCryptoWalletsReply{
		Records:     records,
		Total:       uint64(total),
		CurrentPage: uint32(offset/limit + 1),
	}, nil
}
