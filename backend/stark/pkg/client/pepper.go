// Code generated by night-kit. DO NOT EDIT.
// Version <no value>

package client

import (
	"gitlab.com/greyhole/night-kit/pkg/carbon"
	nightkitgrpc "gitlab.com/greyhole/night-kit/pkg/grpc"
	proto "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"google.golang.org/grpc"
)

type PepperClient struct {
	proto.PepperClient
}

func NewPepperClient(config *api.TCPSocket, options ...grpc.DialOption) (*PepperClient, error) {
	conn, err := nightkitgrpc.NewClientConnection(config, options...)
	if err != nil {
		return nil, err
	}
	return &PepperClient{
		PepperClient: proto.NewPepperClient(conn),
	}, nil
}
