package natasha

import (
	"context"
	"strconv"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	emerchant "gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/merchant"
)

const (
	_defaultLimit  = 10
	_defaultOffset = 0
)

func (s *natashaServer) ListMerchants(ctx context.Context, request *natasha.ListMerchantsRequest) (*natasha.ListMerchantsReply, error) {
	var (
		limit  = _defaultLimit
		offset = _defaultOffset
	)
	if request.GetSize() > 0 {
		limit = int(request.GetSize())
	}
	if request.GetPage() > 0 {
		offset = (int(request.GetPage()) - 1) * limit
	}
	idOrName, _ := strconv.Atoi(request.GetKeyword())

	mexs, err := s.eclient.Merchant.
		Query().
		Where(
			emerchant.Or(
				emerchant.IDEQ(int64(idOrName)),
				emerchant.NameContainsFold(request.GetKeyword()),
			),
		).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}
	all, err := s.eclient.Merchant.Query().All(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}
	records := []*natasha.Merchant{}
	for _, mex := range mexs {
		records = append(records, transformer.GetMerchant(mex))
	}

	return &natasha.ListMerchantsReply{
		Records:     records,
		Total:       uint64(len(all)),
		CurrentPage: uint32(offset/limit + 1),
		CurrentSize: uint32(limit),
	}, nil
}
