package tony

import (
	"context"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
)

const (
	_accountNameNonSpecialRegex = "^[a-zA-Z0-9_]+( [a-zA-Z0-9_]+)*$"
	_accountNumberOnlyNumber    = "^[0-9]+$"
)

func (s *tonyServer) ValidateSystemEWallets(
	ctx context.Context,
	request *stark.ValidateSystemEWalletsRequest,
) (*stark.ValidateSystemEWalletsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	_, _, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	var validRecords []*stark.CreateSystemEWalletRequest
	var duplicatedRecords []*stark.CreateSystemEWalletRequest
	var duplicatedIDRecords []*stark.CreateSystemEWalletRequest
	var invalidRecords []*stark.CreateSystemEWalletRequest
	for _, r := range request.GetRecords() {
		code := s.validateCreateSystemEWallet(ctx, r)
		if code == codes.Code_OK {
			validRecords = append(validRecords, r)
			continue
		}
		if code == codes.Code_EWALLET_ACCOUNT_ID_EXISTED {
			duplicatedIDRecords = append(duplicatedIDRecords, r)
			continue
		}
		if code == codes.Code_EWALLET_ACCOUNT_EXISTED {
			duplicatedRecords = append(duplicatedRecords, r)
			continue
		}
		r.ErrorCode = code.String()
		invalidRecords = append(invalidRecords, r)
	}

	return &stark.ValidateSystemEWalletsReply{
		ValidRecords:        validRecords,
		DuplicatedRecords:   duplicatedRecords,
		InvalidRecords:      invalidRecords,
		DuplicatedIdRecords: duplicatedIDRecords,
	}, nil
}
