package jarvis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

type CallbackCompletePayment func(*CompletePaymentRequest, error)

type CompletePaymentRequest struct {
	PaymentId         int64             `json:"payment_id,omitempty"`
	PaymentMethod     stark.MethodType  `json:"payment_method,omitempty"`
	PaymentType       stark.PaymentType `json:"payment_type,omitempty"`
	PaymentCode       string            `json:"payment_code,omitempty"`
	PaymentStatus     stark.Status      `json:"payment_status,omitempty"`
	MexCurrentBalance uint64            `json:"mex_current_balance,omitempty"`
	PaymentDetail     *PaymentDetail    `json:"payment_detail,omitempty"`
}

type PaymentDetail struct {
	Banking *stark.BankingPaymentDetail `json:"banking,omitempty"`
	EWallet *stark.EWalletPaymentDetail `json:"e_wallet,omitempty"`
	Telco   *stark.TelcoPaymentDetail   `json:"telco,omitempty"`
	Crypto  *stark.CryptoPaymentDetail  `json:"crypto,omitempty"`
}

type Client interface {
	// return error for future queue impelemtation
	CompletePayment(
		ctx context.Context,
		url string,
		request *CompletePaymentRequest,
		callback CallbackCompletePayment,
	) error
}

func New(logger *zap.Logger) Client {
	return &client{
		logger:     logger,
		httpClient: http.Client{},
	}
}

type client struct {
	httpClient http.Client
	logger     *zap.Logger
}

func (c *client) CompletePayment(
	ctx context.Context,
	url string,
	request *CompletePaymentRequest,
	callback CallbackCompletePayment,
) error {

	go func() {
		// Calling sleep method to simulate latency
		time.Sleep(3 * time.Second)
		body, err := json.Marshal(request)
		if err != nil {
			callback(request, err)
			return
		}
		c.logger.Debug("complete payment", zap.Any("body ", string(body)))
		resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			callback(request, err)
			return
		}
		c.logger.Debug("response from mex", zap.Any("resp", resp))
		if resp.StatusCode != http.StatusOK {
			callback(request, fmt.Errorf("call to mex failed with code %s", resp.Status))
			return
		}
		callback(request, nil)
	}()
	return nil
}
