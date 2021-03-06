// Code generated by night-kit. DO NOT EDIT.
// Version <no value>

package client

import (
	"gitlab.com/greyhole/night-kit/pkg/carbon"
	nightkitgrpc "gitlab.com/greyhole/night-kit/pkg/grpc"
	proto "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
	"google.golang.org/grpc"
)

type VisionClient struct {
	proto.VisionClient
}

func NewVisionClient(config *api.TCPSocket, options ...grpc.DialOption) (*VisionClient, error) {
	conn, err := nightkitgrpc.NewClientConnection(config, options...)
	if err != nil {
		return nil, err
	}
	return &VisionClient{
		VisionClient: proto.NewVisionClient(conn),
	}, nil
}
