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

type napTuDongParams struct {
	Status        int64  `schema:"status"`
	Message       string `schema:"message"`
	RequestID     int64  `schema:"request_id"`
	DeclaredValue string `schema:"declared_value"`
	Value         uint64 `schema:"value"`
	Amount        uint64 `schema:"amount"`
	Code          string `schema:"code"`
	Serial        string `schema:"serial"`
	Telco         string `schema:"telco"`
	TransID       string `schema:"trans_id"`
	CallbackSign  string `schema:"callback_sign"`
}

func (s *httpServer) getNapTuDong(w http.ResponseWriter, r *http.Request) {
	logger := s.logger
	logger.Info("Callback thu the cao")

	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := new(napTuDongParams)
	if err := schema.NewDecoder().Decode(params, r.Form); err != nil {
		logger.Error("failed to decode", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("params ", zap.Any("params", params))
	md := metadata.Pairs("x-user-id", "1")
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	if params.Status != 1 {
		logger.Error("failed to charge telco card", zap.Any("params", params))
		_, err := s.morganClient.RejectTelcoTopUp(ctx, &stark.RejectTelcoTopUpRequest{
			PaymentId: params.RequestID,
			Note: params.Message,
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
