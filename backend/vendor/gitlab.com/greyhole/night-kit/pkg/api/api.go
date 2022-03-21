package api

import (
	"context"
)

// Meta represents static endpoint metadata
type Meta struct {
	Name        string
	Type        string
	Client      string
	Service     string
	CircuitName string
	Idempotent  bool
}

// Endpoint is the basic endpoint type
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Middleware is the basic middleware type
type Middleware func(Meta, Endpoint) Endpoint

func init() { return }
