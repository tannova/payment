package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"net"
	"sync"
)

type Service interface {
	Logger() *zap.Logger
	Stats() *statsd.Client
	Listener() net.Listener
	Server() *grpc.Server
	Init(opts ...Option)
	Serve()
	Options() Options
	MetricRegistry() *prometheus.Registry
	HealthServer() *health.Server
}

type service struct {
	opts Options
}

func (s *service) Logger() *zap.Logger {
	return s.opts.Logger
}

func (s *service) Stats() *statsd.Client {
	return s.opts.Stats
}

func (s *service) Listener() net.Listener {
	return s.opts.Listener
}

func (s *service) Server() *grpc.Server {
	return s.opts.Server
}

func (s *service) HealthServer() *health.Server {
	return s.opts.HealthServer
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) Serve() {
	logger := s.opts.Logger

	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			logger.Fatal("failed to exec before start", zap.Error(err))
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)

	s.serve(wg)

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			logger.Fatal("failed to exec before start", zap.Error(err))
		}
	}

	wg.Wait()
}

func (s *service) MetricRegistry() *prometheus.Registry {
	return s.opts.MetricRegistry
}

func (s *service) serve(wg sync.WaitGroup) {
	logger := s.opts.Logger
	listener := s.opts.Listener
	go func() {
		defer wg.Done()
		logger.Info("Listening", zap.String("address", listener.Addr().String()))
		if err := s.Server().Serve(listener); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()
}

func (s *service) Init(opts ...Option) {
	for _, opt := range opts {
		opt(&s.opts)
	}
}

func NewService(opts ...Option) Service {
	o := newOptions()

	for _, opt := range opts {
		opt(&o)
	}

	s := &service{
		opts: o,
	}

	return s
}
