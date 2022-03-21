package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/health"
	"net"

	"github.com/DataDog/datadog-go/statsd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Options struct {
	Name           string
	Logger         *zap.Logger
	Stats          *statsd.Client
	Server         *grpc.Server
	Listener       net.Listener
	MetricRegistry *prometheus.Registry
	HealthServer   *health.Server

	BeforeStart []func() error
	AfterStart  []func() error
}

type Option func(o *Options)

func Logger(logger *zap.Logger) Option {
	return func(o *Options) {
		o.Logger = logger
	}
}

func Stats(stats *statsd.Client) Option {
	return func(o *Options) {
		o.Stats = stats
	}
}

func Listener(listener net.Listener) Option {
	return func(o *Options) {
		o.Listener = listener
	}
}

func Server(server *grpc.Server) Option {
	return func(o *Options) {
		o.Server = server
	}
}

func HealthServer(healthServer *health.Server) Option {
	return func(o *Options) {
		o.HealthServer = healthServer
	}
}

func MetricRegistry(registry *prometheus.Registry) Option {
	return func(o *Options) {
		o.MetricRegistry = registry
	}
}

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func newOptions() Options {
	return Options{}
}
