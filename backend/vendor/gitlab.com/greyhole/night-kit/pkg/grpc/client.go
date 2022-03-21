package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	api "gitlab.com/greyhole/night-kit/pkg/carbon"
	"google.golang.org/grpc/credentials"
)

func NewClientConnection(config *api.TCPSocket, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(InjectRequestMetadata),
	}
	if config.GetSecure() {
		options = append(options, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	} else {
		options = append(options, grpc.WithInsecure())
	}
	options = append(options, opts...)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Address, config.GetPort()), options...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func InjectRequestMetadata(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	inMd, ok := metadata.FromIncomingContext(ctx)
	if ok {
		requestId := inMd.Get(XRequestIDHeader)
		md.Set(XRequestIDHeader, requestId...)
	}

	newCtx := metadata.NewOutgoingContext(ctx, md)
	return invoker(newCtx, method, req, reply, cc, opts...)
}
