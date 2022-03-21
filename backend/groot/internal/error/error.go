package error

import (
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	code "gitlab.com/mcuc/monorepo/backend/groot/pkg/code"
)

var (
	NotFound = Error(code.Code_NOT_FOUND)
)

func New(c code.Code, msg string) error {
	code := grpccodes.Code(c)
	return grpcstatus.New(code, msg).Err()
}

func Error(c code.Code) error {
	return New(c, c.String())
}

func Internal(msg string) error {
	return New(code.Code_INTERNAL, msg)
}

