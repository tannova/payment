package morgan

import (
	"context"
	"encoding/json"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *morganServer) UpdateGetCardProvidersSetting(ctx context.Context, request *stark.UpdateGetCardProvidersSettingRequest) (*stark.UpdateGetCardProvidersSettingReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	arr, err := json.Marshal(request.Providers)
	if err != nil {
		logger.Error("failed to marshal", zap.Error(err))
		return nil, status.Internal(err.Error())
	}
	s.entClient.Setting.
		Create().
		SetName("get_card_third_party_priorities").
		SetType("array").
		SetValue(string(arr)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	return &stark.UpdateGetCardProvidersSettingReply{}, nil
}
