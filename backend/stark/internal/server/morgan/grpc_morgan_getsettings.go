package morgan

import (
	"context"

	"gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

func (s *morganServer) GetSettings(ctx context.Context, request *stark.GetSettingsRequest) (*stark.GetSettingsReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.telcoSetting.GetSettings(ctx, request)
}
