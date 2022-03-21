package logging

import (
	"context"
	"fmt"
	"gitlab.com/greyhole/night-kit/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"gitlab.com/greyhole/night-kit/pkg/carbon"
)

// NewLogger ...
func NewLogger(msg *api.Logger) (*zap.Logger, error) {
	var c zap.Config
	var opts []zap.Option
	if msg.GetPretty() {
		c = zap.NewDevelopmentConfig()
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	} else {
		c = zap.NewProductionConfig()
	}

	level := zap.NewAtomicLevel()

	levelName := "INFO"
	if msg.Level != api.Logger_UNSPECIFIED {
		levelName = msg.Level.String()
	}

	if err := level.UnmarshalText([]byte(levelName)); err != nil {
		return nil, fmt.Errorf("could not parse log level %s", msg.Level.String())
	}
	c.Level = level

	return c.Build(opts...)
}

var _logger *zap.Logger = nil

func InitLogger(msg *api.Logger) (err error) {
	_logger, err = NewLogger(msg)
	return err
}

// NewTmpLogger ...
func NewTmpLogger() *zap.Logger {
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	return l
}

// Logger Return new logger with context value
// ctx:  nullable
func Logger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return _logger
	}

	requestID := getRequestID(ctx)
	return _logger.With(zap.String(grpc.XRequestIDHeader, requestID))
}

func getRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	requestIds := md.Get(grpc.XRequestIDHeader)
	if len(requestIds) < 1 {
		return ""
	}
	return requestIds[0]
}
