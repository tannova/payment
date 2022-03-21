package banthe247

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"

	ie "gitlab.com/mcuc/monorepo/backend/groot/internal/error"
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	codes "gitlab.com/mcuc/monorepo/backend/groot/pkg/code"
)

const (
	_getCardPath = "/v2/PayCards/TelcoPay/GetCards"
)

var (
	mapTelcoName = map[groot.TelcoName]string {
		groot.TelcoName_VIETTEL: "VTT",
		groot.TelcoName_MOBIPHONE: "VMS",
		groot.TelcoName_VIETNAMMOBILE: "VNM",
		groot.TelcoName_VINAPHONE: "VNP",
	}

	mapStringTelcoName = map[string]groot.TelcoName {
		"VTT": groot.TelcoName_VIETTEL,
		"VMS": groot.TelcoName_MOBIPHONE,
		"VNM": groot.TelcoName_VIETNAMMOBILE,
		"VNP": groot.TelcoName_VINAPHONE,
	}
)

type GetCardRequest struct {
	Token string
	TelcoName groot.TelcoName
	Amount uint64
}

type GetCardReply struct {
	Card groot.Card
}

type getCardReply struct {
	ErrorCode int64 `json:"errorCode"`
	Message string `json:"message"`
	Data string `json:"Data"`
}

type cardDetail struct {
	PinCode string `json:"PinCode"`
	Telco   string `json:"Telco"`
	Serial  string `json:"Serial"`
	Amount  json.Number `json:"Amount"`
	Trace   json.Number `json:"Trace"`
}

type cardInfo []cardDetail

func (p provider) GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	logger := logging.Logger(ctx)

	tokenResp, err := p.GetToken(ctx)
	if err != nil {
		logger.Error("failed to get token", zap.Error(err))
		return nil, err
	}

	if tokenResp.Token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	r, err := http.NewRequest("POST", p.cfg.baseURL + _getCardPath, nil)
	if err != nil {
		logger.Error("failed to init request", zap.Error(err))
		return nil, err
	}

	client := &http.Client{}
	q := r.URL.Query()
	q.Add("msg", p.buildGetCardMsg(request.TelcoName, request.Amount))
	decoded, err := url.QueryUnescape(q.Encode())
	r.URL.RawQuery = decoded
	r.WithContext(ctx)
	r.Header.Set("Token", tokenResp.Token)

	resp, err := client.Do(r)
	if err != nil {
		logger.Error("failed to create get card request", zap.Error(err))
		return nil, err
	}
	resJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", zap.Error(err))
		return nil, err
	}

	_, err = p.repository.Transaction.Create(ctx, &ent.Transaction{
		Provider: _providerName,
		Amount: request.Amount,
		Request: r.URL.String(),
		Response: string(resJson),
		TxID: uuid.New().String(),
	})
	if err != nil {
		logger.Error("failed to store transaction", zap.Error(err))
		return nil, err
	}
	var res getCardReply

	if err := json.Unmarshal(resJson, &res); err != nil {
		logger.Error("failed to unmarshal json response", zap.Error(err), zap.String("response", string(resJson)))
		return nil, ie.Internal(err.Error())
	}

	if res.ErrorCode == 778 {
		return nil, ie.Error(codes.Code_INSUFFICIENT_BALANCE)
	}

	if res.ErrorCode != 0 {
		return nil, ie.Internal(res.Message)
	}

	var card cardInfo
	if err := json.Unmarshal([]byte(res.Data), &card); err != nil {
		logger.Error("failed to unmarshal json response", zap.Error(err), zap.String("response", res.Data))
		return nil, ie.Internal(err.Error())
	}
	amount, _ := card[0].Amount.Int64()

	return &groot.Card{
		Code: card[0].PinCode,
		Serial: card[0].Serial,
		Amount: uint64(amount),
		TelcoName: mapStringTelcoName[card[0].Telco],
	}, nil
}

func (p provider) buildGetCardMsg(branch groot.TelcoName, amount uint64) string {
	return fmt.Sprintf("%s:%d:1", mapTelcoName[branch], amount)
}