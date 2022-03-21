package natasha

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

func (s *natashaServer) UpdateMerchant(ctx context.Context, request *natasha.UpdateMerchantRequest) (*natasha.UpdateMerchantReply, error) {
	if err := request.Validate(); err != nil {
		s.logger.Warn("invalid request", zap.Error(err))
		// TODO(tiennv147): define error code
		return nil, err
	}
	_, tellerID, err := getUserID(ctx)
	if err != nil {
		s.logger.Error("can not get teller id", zap.Error(err))
		return nil, err
	}

	updateOne := s.eclient.Merchant.UpdateOneID(request.GetId())
	updateOne.SetUpdatedBy(tellerID)
	if request.GetName() != "" {
		updateOne.SetName(request.GetName())
	}
	if request.GetLogoPath() != "" {
		updateOne.SetLogoPath(request.GetLogoPath())
	}
	if request.GetWebhookUrl() != "" {
		updateOne.SetWebhookURL(request.GetWebhookUrl())
	}
	if request.GetEmailContact() != "" {
		updateOne.SetEmailContact(request.GetEmailContact())
	}
	updateOne.SetSafetyLimit(request.GetSafetyLimit())

	mex, err := updateOne.Save(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		return nil, err
	}

	return &natasha.UpdateMerchantReply{
		Merchant: transformer.GetMerchant(mex),
	}, nil
}
