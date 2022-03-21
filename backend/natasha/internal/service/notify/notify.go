package notify

import (
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

type Notify interface {
	Send(msg *slack.WebhookMessage, slackWebhookUrl string)
}

type notify struct {
	logger *zap.Logger
}

func NewNotify(logger *zap.Logger) Notify {
	return &notify{
		logger: logger,
	}
}

func (n notify) Send(msg *slack.WebhookMessage, slackWebhookUrl string) {
	if slackWebhookUrl == "" {
		n.logger.Warn("Webhook URL empty")
		return
	}
	if msg == nil {
		n.logger.Warn("Message is empty")
		return
	}

	err := slack.PostWebhook(slackWebhookUrl, msg)
	if err != nil {
		n.logger.Error("Post web hook failed", zap.Error(err))
	}
}
