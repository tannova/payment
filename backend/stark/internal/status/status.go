package status

import (
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	code "gitlab.com/mcuc/monorepo/backend/stark/pkg/code"
)

var (
	Unauthenticated = Error(code.Code_UNAUTHENTICATED)
)

func New(c code.Code, msg string) error {
	grpcCode := grpccodes.Code(c)
	return grpcstatus.New(grpcCode, msg).Err()
}

func Error(c code.Code) error {
	return New(c, c.String())
}

func Internal(msg string) error {
	return New(code.Code_INTERNAL, msg)
}

func InvalidArgument(msg string) error {
	return New(code.Code_INVALID_ARGUMENT, msg)
}

func Code(err error) code.Code {
	return code.Code(grpcstatus.Code(err))
}
