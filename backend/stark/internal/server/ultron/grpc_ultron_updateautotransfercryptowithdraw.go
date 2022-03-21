package ultron

import (
	"context"
	"fmt"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

const (
	_autoTransferWithdrawCryptoName = "auto_transfer_withdraw_crypto"
)

func (s *ultronServer) UpdateAutoTransferCryptoWithdraw(
	ctx context.Context,
	request *stark.UpdateAutoTransferCryptoWithdrawRequest,
) (*stark.UpdateAutoTransferCryptoWithdrawReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	err := s.entClient.Setting.
		Create().
		SetName(_autoTransferWithdrawCryptoName).
		SetType("bool").
		SetValue(fmt.Sprint(request.Enabled)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		logger.Error("update setting err", zap.Error(err))
		return nil, err
	}

	return &stark.UpdateAutoTransferCryptoWithdrawReply{}, nil
}
