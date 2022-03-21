package blackwidow

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"go.uber.org/zap"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	postman "gitlab.com/greyhole/postman/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
	templateLookup "gitlab.com/mcuc/monorepo/backend/natasha/template"
)

const (
	_emailSubject    = "AloPay: The payment is rejected"
	_contentTemplate = "PaymentID: %d, reason: %s"
)

func (s *blackWidowServer) NotifyRejectPayment(ctx context.Context, request *natasha.NotifyRejectPaymentRequest) (*natasha.NotifyRejectPaymentReply, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	logger := logging.Logger(ctx)
	merchant, err := s.repository.Merchant.GetMerchant(ctx, request.GetMerchantId())
	if err != nil {
		logger.Error("failed to get merchant", zap.Error(err))
		return nil, err
	}
	params := map[string]interface{}{
		"PaymentID": request.PaymentId,
	}
	emailTemplate := templateLookup.Lookup(s.cfg.MerchantCall.GetEmail().GetTemplate())
	content, err := createContent(logger, emailTemplate, params)
	if err != nil {
		logger.Error("failed to get email content", zap.Error(err))
		return nil, err
	}
	emailRequest := &postman.SendMailRequest{
		Recipients: []string{merchant.EmailContact},
		Subject:    s.cfg.MerchantCall.Email.Subject,
		Content:    fmt.Sprintf(content, request.PaymentId, request.Reason),
	}

	if _, err := s.postman.SendMail(ctx, emailRequest); err != nil {
		s.logger.Error("can not send email", zap.Error(err))
		return nil, err
	}

	return &natasha.NotifyRejectPaymentReply{}, nil
}

func createContent(logger *zap.Logger, template *template.Template, data interface{}) (string, error) {
	var b bytes.Buffer
	if err := template.Execute(&b, data); err != nil {
		logger.Error("can not execute template", zap.Error(err))
		return "", err
	}

	return b.String(), nil
}
