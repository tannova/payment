package utils

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strconv"

	"gitlab.com/greyhole/myid/pkg/extractor"
)

func GetUserIDFromContext(ctx context.Context) (int64, string, error) {
	n, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return 0, "", err
	}
	return n, strconv.FormatInt(n, 10), nil
}

func GetMetadata(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := errors.New("no metadata from context")
		return nil, err
	}

	return md, nil
}
