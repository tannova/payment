package ultron

import (
	"context"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/stark/internal/status"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent"
)

func (s *ultronServer) GetCryptoSettings(
	ctx context.Context,
	request *stark.GetCryptoSettingsRequest,
) (*stark.GetCryptoSettingsReply, error) {

	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)

	settings, err := s.entClient.Setting.Query().All(ctx)
	if err != nil {
		logger.Error("failed to get setting", zap.Error(err))
		return nil, status.Internal(err.Error())
	}

	m := make(map[string]*ent.Setting)

	for _, s := range settings {
		m[s.Name] = s
	}

	autoTransferWithdrawCrypto := false
	if s, ok := m[_autoTransferWithdrawCryptoName]; ok {
		autoTransferWithdrawCrypto = s.Value == "true"
	}

	return &stark.GetCryptoSettingsReply{
		AutoTransferWithdrawCrypto: autoTransferWithdrawCrypto,
	}, nil
}
