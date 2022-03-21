package bishop

import (
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"net/http"

	"github.com/gorilla/schema"
	"go.uber.org/zap"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

type theNhanh365Params struct {
	RequestID     int64  `schema:"request_id"`
	Value         uint64 `schema:"value"`
	Status        int64  `schema:"status"`
	Amount        uint64 `schema:"amount"`
	Serial        string `schema:"serial"`
	Telco         string `schema:"telco"`
	Code          string `schema:"code"`
	DeclaredValue uint64 `schema:"declared_value"`
	CallbackSign  string `schema:"callback_sign"`
	TransId       string `schema:"trans_id"`
}

var (
	mapTheNhanh365TelcoName = map[string]groot.TelcoName{
		"VIETTEL":      groot.TelcoName_VIETTEL,
		"MOBIFONE":     groot.TelcoName_MOBIPHONE,
		"VIETNAMOBILE": groot.TelcoName_VIETNAMMOBILE,
		"VINAPHONE":    groot.TelcoName_VINAPHONE,
	}
)

func (s *httpServer) getTheNhanh365(w http.ResponseWriter, r *http.Request) {
	logger := s.logger
	logger.Info("Callback thu the cao")

	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := new(theNhanh365Params)
	if err := schema.NewDecoder().Decode(params, r.Form); err != nil {
		logger.Error("failed to decode", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("params ", zap.Any("params", params))
	md := metadata.Pairs("x-user-id", "1")
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	if params.Status != 1 {
		logger.Error("failed to verify card", zap.Any("params", params))
		_, err := s.morganClient.RejectTelcoTopUp(ctx, &stark.RejectTelcoTopUpRequest{
			PaymentId: params.RequestID,
			Note: "Invalid Card",
		})
		if err != nil {
			logger.Error("failed to reject payment", zap.Error(err))
			return
		}
		return
	}

	_, err := s.morganClient.CompleteTelcoTopUp(ctx, &stark.CompleteTelcoTopUpRequest{
		Card: &groot.Card{
			Amount:    params.Value,
			Serial:    params.Serial,
			TelcoName: mapTheNhanh365TelcoName[params.Telco],
			Code:      params.Code,
		},
		PaymentId:     params.RequestID,
		ChargedAmount: params.Amount,
	})
	if err != nil {
		logger.Error("failed to complete payment", zap.Error(err))
		return
	}

	// create response
	data := make(map[string]string)
	data["message"] = "OK"
	response := Response{
		Success: true,
	}

	// make json
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content-type and return
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
