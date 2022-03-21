package tony

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
	esw "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
)

func (s *tonyServer) ListSystemEWallets(ctx context.Context, request *stark.ListSystemEWalletsRequest) (*stark.ListSystemEWalletsReply, error) {
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

	var filters []predicate.SystemEWallet
	ids := request.GetIds()
	if len(ids) > 0 {
		filters = append(filters, esw.IDIn(ids...))
	}

	merchantIDs := request.GetMerchantIds()
	if len(merchantIDs) > 0 {
		filters = append(filters, esw.MerchantIDIn(merchantIDs...))
	}

	statuses := request.GetStatuses()
	if len(statuses) > 0 {
		in := []int32{}
		for _, v := range statuses {
			in = append(in, int32(v))
		}
		filters = append(filters, esw.StatusIn(in...))
	}

	eWalletNames := request.GetEWalletNames()
	if len(eWalletNames) > 0 {
		in := []int32{}
		for _, v := range eWalletNames {
			in = append(in, int32(v))
		}
		filters = append(filters, esw.EWalletNameIn(in...))
	}

	wallets, err := s.entClient.SystemEWallet.
		Query().
		Where(esw.And(filters...)).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(esw.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		logger.Error("could not query crypto wallets", zap.Error(err))
		return nil, err
	}
	total, err := s.entClient.SystemEWallet.
		Query().
		Where(esw.And(filters...)).
		Count(ctx)
	if err != nil {
		logger.Error("could not count crypto wallets", zap.Error(err))
		return nil, err
	}
	records := []*stark.SystemEWallet{}
	for _, wallet := range wallets {
		records = append(records, transformer.GetSystemEWallet(wallet))
	}

	return &stark.ListSystemEWalletsReply{
		Records:     records,
		Total:       uint64(total),
		CurrentPage: uint32(offset/limit + 1),
	}, nil
}
