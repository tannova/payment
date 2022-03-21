package pepper

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/transformer"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
	sba "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

func (s *pepperServer) ListSystemBankAccounts(ctx context.Context, request *stark.ListSystemBankAccountsRequest) (*stark.ListSystemBankAccountsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	logger := logging.Logger(ctx)

	var (
		limit  = 10
		offset = 0
	)

	_, _, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Unauthenticated
	}

	if request.GetSize() > 0 {
		limit = int(request.GetSize())
	}
	if request.GetPage() > 0 {
		offset = (int(request.GetPage()) - 1) * limit
	}

	var filters []predicate.SystemBankAccount
	if len(request.GetStatues()) > 0 {
		statuses := make([]int32, len(request.GetStatues()))
		for i, v := range request.GetStatues() {
			statuses[i] = int32(v)
		}
		filters = append(filters, sba.StatusIn(statuses...))
	}

	merchantIDs := request.GetMerchantIds()
	if len(merchantIDs) > 0 {
		filters = append(filters, sba.MerchantIDIn(merchantIDs...))
	}

	if len(request.GetBankNames()) > 0 {
		bankNames := make([]int32, len(request.GetBankNames()))
		for i, v := range request.GetBankNames() {
			bankNames[i] = int32(v)
		}
		filters = append(filters, sba.BankNameIn(bankNames...))
	}

	if len(request.GetIds()) > 0 {
		filters = append(filters, sba.IDIn(request.GetIds()...))
	}

	accounts, err := s.entClient.SystemBankAccount.
		Query().
		Where(sba.And(filters...)).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(sba.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		logger.Error("failed to query list system bank account", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	total, err := s.entClient.SystemBankAccount.
		Query().
		Where(sba.And(filters...)).
		Count(ctx)
	if err != nil {
		logger.Error("failed to count system bank account", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	records := make([]*stark.SystemBank, len(accounts))
	for i, a := range accounts {
		records[i] = transformer.GetSystemBankAccount(a)
	}

	return &stark.ListSystemBankAccountsReply{
		Records:     records,
		Total:       int64(total),
		CurrentPage: uint32(offset/limit + 1),
	}, nil
}
