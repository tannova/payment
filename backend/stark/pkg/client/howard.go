// Code generated by night-kit. DO NOT EDIT.
// Version <no value>

package client

import (
	"gitlab.com/greyhole/night-kit/pkg/carbon"
	nightkitgrpc "gitlab.com/greyhole/night-kit/pkg/grpc"
	proto "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"google.golang.org/grpc"
)

type HowardClient struct {
	proto.HowardClient
}

func NewHowardClient(config *api.TCPSocket, options ...grpc.DialOption) (*HowardClient, error) {
	conn, err := nightkitgrpc.NewClientConnection(config, options...)
	if err != nil {
		return nil, err
	}
	return &HowardClient{
		HowardClient: proto.NewHowardClient(conn),
	}, nil
}
