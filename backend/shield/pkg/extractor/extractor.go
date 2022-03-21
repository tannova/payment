package extractor

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	// MexID is merchant ID
	// MexIDHeader = "x-mex-id"
	// TellerID is teller ID
	// TellerIDHeader = "x-teller-id"

	// Temporary
	MexIDHeader    = "x-cmgp-userid"
	TellerIDHeader = "x-cmgp-userid"
)

var (
	// ErrNoMetadata ...
	ErrNoMetadata = errors.New("no metadata from context")
	// ErrEmptyMetadata ...
	ErrEmptyMetadata = errors.New("empty metadata")
)

// Extractor put it to an interface for easier when we want to mock
type Extractor interface {
	GetMexID(ctx context.Context) (int64, error)
	GetTellerID(ctx context.Context) (int64, error)
}

func NewExtractor() Extractor {
	return &extractorImpl{}
}

type extractorImpl struct {
}

// GetMexID get merchant ID
func (i *extractorImpl) GetMexID(ctx context.Context) (int64, error) {
	return getMetadata(ctx, MexIDHeader, 0)
}

// GetTellerID ...
func (i *extractorImpl) GetTellerID(ctx context.Context) (int64, error) {
	return getMetadata(ctx, TellerIDHeader, 0)
}

func getMetadata(ctx context.Context, key string, defval int64) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return defval, ErrNoMetadata
	}
	mdVal := md.Get(key)
	if len(mdVal) < 1 {
		return defval, ErrEmptyMetadata
	}

	id, err := strconv.ParseInt(mdVal[0], 10, 64)
	if err != nil {
		return defval, ErrEmptyMetadata
	}

	return id, nil
}
