package thuthecao

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.com/greyhole/night-kit/pkg/logging"
	ie "gitlab.com/mcuc/monorepo/backend/groot/internal/error"
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_chargePath = "/api/card-auto.php"
)

var (
	mapTelcoName = map[groot.TelcoName]string {
		groot.TelcoName_VIETTEL: "VIETTEL",
		groot.TelcoName_MOBIPHONE: "MOBIFONE",
		groot.TelcoName_VIETNAMMOBILE: "VNMOBI",
		groot.TelcoName_VINAPHONE: "VINAPHONE",
	}
)

type chargeCardReply struct {
	Data struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	} `json:"data"`
}

func (p provider) GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (p provider) ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) error {
	logger := logging.Logger(ctx)

	r, err := http.NewRequest("GET", p.cfg.baseUrl + _chargePath, nil)
	if err != nil {
		logger.Error("failed to init request", zap.Error(err))
		return err
	}

	client := &http.Client{}
	q := r.URL.Query()

	branch, ok := mapTelcoName[request.Card.TelcoName]
	if !ok {
		return ie.Internal(fmt.Sprintf("branch %d not found", request.Card.TelcoName))
	}
	q.Add("type", branch)
	q.Add("menhgia", fmt.Sprint(request.Card.Amount))
	q.Add("seri", request.Card.Serial)
	q.Add("pin", request.Card.Code)
	q.Add("APIKey", p.cfg.apiKey)
	q.Add("callback", p.cfg.callbackUrl)
	q.Add("content", request.PaymentId)
	decoded, err := url.QueryUnescape(q.Encode())
	r.URL.RawQuery = decoded
	r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		logger.Error("failed to get token", zap.Error(err))
		return err
	}

	resJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", zap.Error(err))
		return err
	}

	_, err = p.repository.Transaction.Create(ctx, &ent.Transaction{
		Provider: _providerName,
		Amount: request.Card.Amount,
		Request: r.URL.String(),
		Response: string(resJson),
		TxID: request.PaymentId,
		Status: "Pending",
	})
	if err != nil {
		logger.Error("failed to store transaction", zap.Error(err))
		return err
	}

	var res chargeCardReply
	if err := json.Unmarshal(resJson, &res); err != nil {
		logger.Error("failed to unmarshal json response",
						zap.Error(err),
						zap.String("response", string(resJson)),
						zap.String("request", r.URL.String()))
		return ie.Internal(err.Error())
	}

	if res.Data.Status != "success" {
		return ie.Internal(res.Data.Msg)
	}

	return nil
}
