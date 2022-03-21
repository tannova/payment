package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
	pes "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemcryptohotwallet"
)

func (s *ultronServer) ListCryptoHotWallets(
	ctx context.Context,
	request *stark.ListCryptoHotWalletsRequest,
) (*stark.ListCryptoHotWalletsReply, error) {

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

	var filters []predicate.SystemCryptoHotWallet
	addresses := request.GetAddresses()
	if len(addresses) > 0 {
		filters = append(filters, pes.AddressIn(addresses...))
	}

	cryptoTypes := request.GetCryptoTypes()
	if len(cryptoTypes) > 0 {
		in := []int32{}
		for _, v := range cryptoTypes {
			in = append(in, int32(v))
		}
		filters = append(filters, pes.CryptoTypeIn(in...))
	}

	cryptoNetworkTypes := request.GetCryptoNetworkTypes()
	if len(cryptoNetworkTypes) > 0 {
		in := []int32{}
		for _, v := range cryptoNetworkTypes {
			in = append(in, int32(v))
		}
		filters = append(filters, pes.CryptoNetworkTypeIn(in...))
	}

	statuses := request.GetStatuses()
	if len(statuses) > 0 {
		in := []int32{}
		for _, v := range statuses {
			in = append(in, int32(v))
		}
		filters = append(filters, pes.StatusIn(in...))
	}

	merchantIDs := request.GetMerchantIds()
	if len(merchantIDs) > 0 {
		filters = append(filters, pes.MerchantIDIn(merchantIDs...))
	}

	wallets, err := s.entClient.SystemCryptoHotWallet.
		Query().
		Where(pes.And(filters...)).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(pes.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		logger.Error("could not query crypto wallets", zap.Error(err))
		return nil, err
	}
	total, err := s.entClient.SystemCryptoHotWallet.
		Query().
		Where(pes.And(filters...)).
		Count(ctx)
	if err != nil {
		logger.Error("could not count crypto wallets", zap.Error(err))
		return nil, err
	}
	records := []*stark.SystemCryptoHotWallet{}
	for _, wallet := range wallets {
		records = append(records, transformer.GetSystemCryptoHotWallet(wallet))
	}

	return &stark.ListCryptoHotWalletsReply{
		Records:     records,
		Total:       uint64(total),
		CurrentPage: uint32(offset/limit + 1),
	}, nil
}
