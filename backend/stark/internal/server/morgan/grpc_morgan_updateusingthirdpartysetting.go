package morgan

import (
	"context"
	"fmt"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *morganServer) UpdateUsingThirdPartySetting(ctx context.Context, request *stark.UpdateUsingThirdPartySettingRequest) (*stark.UpdateUsingThirdPartySettingReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	s.entClient.Setting.
		Create().
		SetName(_enableThirdPartySettingName).
		SetType("bool").
		SetValue(fmt.Sprint(request.EnableThirdParty)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	return &stark.UpdateUsingThirdPartySettingReply{}, nil
}
