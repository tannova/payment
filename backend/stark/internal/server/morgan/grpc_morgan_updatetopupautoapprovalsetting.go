package morgan

import (
	"context"
	"fmt"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *morganServer) UpdateTopUpAutoApprovalSetting(ctx context.Context, request *stark.UpdateTopUpAutoApprovalSettingRequest) (*stark.UpdateTopUpAutoApprovalSettingReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	s.entClient.Setting.
		Create().
		SetName(_autoApprovalSettingName).
		SetType("bool").
		SetValue(fmt.Sprint(request.TopUpAutoApproval)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	return &stark.UpdateTopUpAutoApprovalSettingReply{}, nil
}
