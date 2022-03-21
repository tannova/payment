package natasha

import (
	"context"
	"fmt"
	"time"

	postman "gitlab.com/greyhole/postman/pkg/api"
	"go.uber.org/zap"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/transformer"
	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

const (
	emailSubject    = "AloPay: New Merchant Details"
	contentTemplate = "Your JWT Token for intergration %s"
)

func (s *natashaServer) CreateMerchant(ctx context.Context, request *natasha.CreateMerchantRequest) (*natasha.CreateMerchantReply, error) {
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

	mex, err := s.eclient.Merchant.
		Create().
		SetCreatedAt(time.Now().UTC()).
		SetName(request.GetName()).
		SetEmailContact(request.GetEmailContact()).
		SetLogoPath(request.GetLogoPath()).
		SetSlackWebhookURL(request.GetSlackWebhookUrl()).
		SetWebhookURL(request.GetWebhookUrl()).
		SetSafetyLimit(request.GetSafetyLimit()).
		SetCreatedBy(tellerID).
		SetUpdatedBy(tellerID).
		Save(ctx)
	if err != nil {
		// TODO(tiennv147): define error code
		s.logger.Error("can not create merchant", zap.Error(err))
		return nil, err
	}

	_, token, err := s.token.Create(mex.ID)
	if err != nil {
		s.logger.Error("can not create token", zap.Error(err))
		return nil, err
	}

	emailRequest := &postman.SendMailRequest{
		Recipients: []string{request.GetEmailContact()},
		Subject:    emailSubject,
		Content:    fmt.Sprintf(contentTemplate, token),
	}

	if _, err := s.postman.SendMail(ctx, emailRequest); err != nil {
		s.logger.Error("can not send email", zap.Error(err))
		return nil, err
	}

	return &natasha.CreateMerchantReply{
		Merchant: transformer.GetMerchant(mex),
	}, nil
}
