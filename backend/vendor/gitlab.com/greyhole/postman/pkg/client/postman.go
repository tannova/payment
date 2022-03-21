package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"gitlab.com/greyhole/night-kit/pkg/carbon"
	proto "gitlab.com/greyhole/postman/pkg/api"
)

type PostmanClient struct {
	proto.PostmanClient
}

func NewPostmanClient(config *api.TCPSocket) (*PostmanClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if config.GetSecure() {
		// set credentials in case needing secure
	} else {
		options = append(options, grpc.WithInsecure())
	}

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Address, config.GetPort()), options...)
	if err != nil {
		return nil, err
	}
	return &PostmanClient{
		PostmanClient: proto.NewPostmanClient(conn),
	}, nil
}
