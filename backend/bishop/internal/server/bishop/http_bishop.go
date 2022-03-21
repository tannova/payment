package bishop

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	conf "gitlab.com/mcuc/monorepo/backend/bishop/pkg/config"
	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type HoseHTTPServer interface {
	Serve() error
}

type httpServer struct {
	logger       *zap.Logger
	cfg          *conf.Config
	morganClient stark.MorganClient
}

func NewHTTPServer(logger *zap.Logger, cfg *conf.Config, morganClient stark.MorganClient) HoseHTTPServer {
	return &httpServer{
		logger:       logger,
		cfg:          cfg,
		morganClient: morganClient,
	}
}

func (s *httpServer) Serve() error {
	logger := s.logger
	address := fmt.Sprintf("%s:%d", s.cfg.Listener.GetTcp().Address, s.cfg.Listener.GetTcp().Port)
	mux := http.NewServeMux()

	// bind handler
	mux.HandleFunc("/thenhanh365", s.getTheNhanh365)
	mux.HandleFunc("/naptudong", s.getNapTuDong)
	mux.HandleFunc("/thuthecao", s.getThuTheCao)

	// listener
	err := http.ListenAndServe(address, mux)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
		return err
	}

	return nil
}
