package createsystembankaccount

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
	ebank "gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/utils"
)

type CreateSystemBankAccount interface {
	Create(ctx context.Context, request *stark.CreateSystemBankAccountRequest) (*stark.CreateSystemBankAccountReply, error)
}

func New(
	entClient *ent.Client,
	natashaCli natasha.NatashaClient,
) CreateSystemBankAccount {

	return &createSystemBankAccount{
		entClient:  entClient,
		natashaCli: natashaCli,
	}
}

type createSystemBankAccount struct {
	entClient  *ent.Client
	natashaCli natasha.NatashaClient
}

func (c createSystemBankAccount) Create(ctx context.Context, request *stark.CreateSystemBankAccountRequest) (*stark.CreateSystemBankAccountReply, error) {
	logger := logging.Logger(ctx)

	if err := request.Validate(); err != nil {
		return nil, err
	}
	_, tellerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Error("failed to get teller id", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	mexRequest := &natasha.GetMerchantRequest{Id: request.GetMerchantId()}
	_, err = c.natashaCli.GetMerchant(ctx, mexRequest)
	if err != nil {
		logger.Error("failed to get merchant", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	bankName, ok := stark.BankName_value[request.BankName]
	if !ok {
		logger.Error("invalid bank name", zap.Error(err))
		return nil, status.InvalidArgument("invalid bank name")
	}

	if request.AccountId == 0 {
		exists, err := c.entClient.SystemBankAccount.Query().Where(ebank.AccountNumber(request.AccountNumber)).Exist(ctx)
		if err != nil {
			logger.Error("failed to check bank account number", zap.Error(err))
			return nil, status.Internal(err.Error())
		}
		if exists {
			logger.Error("bank account number existed", zap.Error(err))
			return nil, status.Error(codes.Code_BANK_EXISTED)
		}
	}

	query := c.entClient.SystemBankAccount.Create().
		SetCreatedBy(tellerID).
		SetUpdatedBy(tellerID).
		SetStatus(int32(stark.BankStatus_BANK_STATUS_IN_ACTIVE)).
		SetMerchantID(request.MerchantId).
		SetAccountName(request.AccountName).
		SetAccountNumber(request.AccountNumber).
		SetBranch(request.Branch).
		SetBalance(request.Balance).
		SetBankName(bankName).
		SetDailyBalanceLimit(request.DailyBalanceLimit)
	if request.AccountId != 0 {
		query.SetID(request.AccountId)
	}

	query.
		OnConflict(
			sql.ConflictColumns(ebank.FieldID),
		).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		logger.Error("failed to create system bank account", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	return &stark.CreateSystemBankAccountReply{}, nil
}
