package bishop

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	groot "gitlab.com/mcuc/monorepo/backend/groot/pkg/api"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

type thuthecaoParams struct {
	PaymentID   int64  `schema:"content"`
	MenhGiaThuc uint64 `schema:"menhgiathuc"`
	Status      string `schema:"status"`
	ThucNhan    uint64 `schema:"thucnhan"`
}

func (s *httpServer) getThuTheCao(w http.ResponseWriter, r *http.Request) {
	logger := s.logger
	logger.Info("Callback thu the cao")

	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := new(thuthecaoParams)
	if err := schema.NewDecoder().Decode(params, r.Form); err != nil {
		logger.Error("failed to decode", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("params ", zap.Any("params", params))

	if params.Status != "hoantat" {
		logger.Error("failed to verify card", zap.Any("params", params))
		return
	}

	md := metadata.Pairs("x-user-id", "1")
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	_, err := s.morganClient.CompleteTelcoTopUp(ctx, &stark.CompleteTelcoTopUpRequest{
		Card: &groot.Card{
			Amount: params.MenhGiaThuc,
		},
		PaymentId:     params.PaymentID,
		ChargedAmount: params.ThucNhan,
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
