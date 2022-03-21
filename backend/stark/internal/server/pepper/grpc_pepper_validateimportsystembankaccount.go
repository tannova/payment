package pepper

import (
	"context"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
)

func (s *pepperServer) ValidateImportSystemBankAccount(ctx context.Context, request *stark.ValidateImportSystemBankAccountRequest) (*stark.ValidateImportSystemBankAccountReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	var dupList []*stark.CreateSystemBankAccountRequest
	var dupIdList []*stark.CreateSystemBankAccountRequest
	var validList []*stark.CreateSystemBankAccountRequest
	var invalidList []*stark.CreateSystemBankAccountRequest
	for _, accountRequest := range request.GetRecords() {
		code := s.validateSystemBankAccount(ctx, accountRequest)
		if code == codes.Code_BANK_ACCOUNT_EXISTED {
			dupList = append(dupList, accountRequest)
			continue
		}
		if code == codes.Code_BANK_ACCOUNT_ID_EXISTED {
			dupIdList = append(dupIdList, accountRequest)
			continue
		}
		if code == codes.Code_OK {
			validList = append(validList, accountRequest)
			continue
		}
		accountRequest.ErrorCode = code.String()
		invalidList = append(invalidList, accountRequest)
	}

	return &stark.ValidateImportSystemBankAccountReply{
		ValidRecords:               validList,
		DuplicatedRecords:          dupList,
		DuplicatedAccountIdRecords: dupIdList,
		InvalidRecords:             invalidList,
	}, nil
}
