package naptudong

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.com/greyhole/night-kit/pkg/logging"
	"go.uber.org/zap"

	ie "gitlab.com/mcuc/monorepo/backend/groot/internal/error"
	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent"
)

const (
	_chargeCardPath = "/chargingws/v2"
	_command = "charging"
	_contentType            = "application/json"
)

var (
	mapTelcoName = map[groot.TelcoName]string {
		groot.TelcoName_VIETTEL: "VIETTEL",
		groot.TelcoName_MOBIPHONE: "MOBIFONE",
		groot.TelcoName_VIETNAMMOBILE: "VIETNAMOBILE",
		groot.TelcoName_VINAPHONE: "VINAPHONE",
	}
)

type chargeCardRequest struct {
	Telco     string `json:"telco"`
	Code      string `json:"code"`
	Serial    string `json:"serial"`
	Amount    string `json:"amount"`
	RequestID string `json:"request_id"`
	PartnerID string `json:"partner_id"`
	Command   string `json:"command"`
	Sign      string `json:"sign"`
}

type chargeCardReply struct {
	TransID       int         `json:"trans_id"`
	RequestID     string      `json:"request_id"`
	Amount        int         `json:"amount"`
	Value         interface{} `json:"value"`
	DeclaredValue int         `json:"declared_value"`
	Telco         string      `json:"telco"`
	Serial        string      `json:"serial"`
	Code          string      `json:"code"`
	Status        int         `json:"status"`
	Message       string      `json:"message"`
}

func (p provider) GetCard(ctx context.Context, request *groot.GetCardRequest) (*groot.Card, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (p provider) ChargeCard(ctx context.Context, request *groot.ChargeCardRequest) error {
	logger := logging.Logger(ctx)

	sign := getMD5(strings.Join([]string{
		p.cfg.partnerKey,
		request.Card.Code,
		_command,
		p.cfg.partnerID,
		request.PaymentId,
		request.Card.Serial,
		mapTelcoName[request.Card.TelcoName],
	},""))

	rq, err := json.Marshal(&chargeCardRequest{
		Telco: mapTelcoName[request.Card.TelcoName],
		Code: request.Card.Code,
		Serial: request.Card.Serial,
		Amount: fmt.Sprint(request.Card.Amount),
		RequestID: request.PaymentId,
		PartnerID: p.cfg.partnerID,
		Command: _command,
		Sign: sign,
	})
	if err != nil {
		logger.Error("failed to convert proto to json ")
		return err
	}

	r, err := http.NewRequest("POST", p.cfg.baseUrl + _chargeCardPath, bytes.NewBuffer(rq))
	if err != nil {
		logger.Error("failed to init request", zap.Error(err))
		return err
	}
	r.Header.Set("Content-Type", _contentType)
	r = r.WithContext(ctx)

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		logger.Error("failed to create charge card request", zap.Error(err))
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
		Request: string(rq),
		Response: string(resJson),
		TxID: request.PaymentId,
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

	if res.Status != 99 {
		return ie.Internal(res.Message)
	}

	return nil

}

func getMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}