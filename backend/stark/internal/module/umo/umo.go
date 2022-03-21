package umo

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Client interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
}

// {
// 	"store_id": 795237177952942219,
// 	"app_time": 1647801551770,
// 	"app_user": "M2_0",
// 	"amount": 0.01,
// 	"sender": "0xcfb83c2a7129752ea6d607ab79047e725297d735",
// 	"recipient": "0x565888b262859c58f84ba5ad4d8036c1d034a85d",
// 	"mac": "F9BBF4577D99BBF48F500720934D8A341927CCF840A8549B906C4DBB0EF6393EF2AD2C8C7285515F570AE59AAD45E0A02272A50CA067B4271CF27D58F41D1519",
// 	"app_trans_id": "678",
// 	"currency": "UNU",
// 	"descripton": "abc",
// 	"sender_memo": "",
// 	"recipient_memo": "",
// 	"items": "",
// 	"pending": false
// }

type CreateOrderRequest struct {
	StoreID       int64   `json:"store_id"`
	AppTransID    string  `json:"app_trans_id"`
	AppTime       int64   `json:"app_time"`
	AppUser       string  `json:"app_user"`
	Amount        float64 `json:"amount"`
	Sender        string  `json:"sender"`
	Recipient     string  `json:"recipient"`
	MAC           string  `json:"mac"`
	Currency      string  `json:"currency"`
	Description   string  `json:"descripton,omitempty"`
	SenderMemo    string  `json:"sender_memo"`
	RecipientMemo string  `json:"recipient_memo"`
	Items         string  `json:"items"`
	Pending         bool  `json:"pending"`
}

type CreateOrderReply struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    OrderDetail `json:"data,omitempty"`
}

type OrderDetail struct {
	TransID    string `json:"trans_id,omitempty"`
	Sender     string `json:"sender,omitempty"`
	SenderMemo string `json:"sender_memo,omitempty"`
	Recipient  string `json:"recipient,omitempty"`
	Currency   string `json:"currency,omitempty"`
	StoreID    string `json:"store_id,omitempty"`
	AppTransID string `json:"app_trans_id,omitempty"`
}

func New(
	logger *zap.Logger,
	host string,
	storeID int64,
	secretKey string,
) Client {
	return &client{
		logger:     logger,
		host:       host,
		storeID:    storeID,
		secretKey:  secretKey,
		httpClient: http.Client{},
	}
}

type client struct {
	httpClient http.Client
	logger     *zap.Logger
	host       string
	storeID    int64
	secretKey  string
}

const (
	createOrderPath = "/payment/order/withdraw"
)

const (
	// store_id|app_trans_id|app_time|app_user|items|currency|amount|sender|sender_memo|recipient|recipient_memo|transaction_fee
	hmacTemplate = "%d|%s|%d|%s|%s|%s|%s|%s|%s|%s|%s|null"
)

func (c *client) CreateOrder(
	ctx context.Context,
	request *CreateOrderRequest,
) (*CreateOrderReply, error) {
	// Set store info
	request.StoreID = c.storeID
	request.AppTime = time.Now().UnixNano() / int64(time.Millisecond)

	//var data_input = store_id.concat("|",app_trans_id, "|", app_time, "|", app_user, "|", items, "|", currency, "|", amount, "|",  sender,  "|", sender_memo, "|", recipient, "|", recipient_memo, "|", transaction_fee);

	// do this to format float value without trailing zeros like 0.01000 which is invalid for UMO request
	amountStr := strconv.FormatFloat(request.Amount, 'f', -1, 64)

	// calculate HMAC
	hmacData := fmt.Sprintf(hmacTemplate,
		request.StoreID,
		request.AppTransID,
		request.AppTime,
		request.AppUser,
		request.Items,
		request.Currency,
		amountStr,
		request.Sender,
		request.SenderMemo,
		request.Recipient,
		request.RecipientMemo,
	)
	h := hmac.New(sha512.New, []byte(c.secretKey))
	h.Write([]byte(hmacData))
	hmac := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	request.MAC = hmac

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	c.logger.Info(
		"create order",
		zap.Any("secret_key", c.secretKey),
		zap.Any("hmac", hmac),
		zap.Any("hmacData", hmacData),
		zap.Any("request", string(body)),
	)
	
	url := c.host + createOrderPath
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("call to umo failed with code %s", resp.Status)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	reply := &CreateOrderReply{}
	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
