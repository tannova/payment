package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	config "gitlab.com/greyhole/night-kit/pkg/carbon"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

// NewClient constructor using from other service
func NewClientNatasha(config *config.TCPSocket) (natasha.NatashaClient, error) {
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

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.GetAddress(), config.GetPort()), options...)
	if err != nil {
		return nil, err
	}

	return natasha.NewNatashaClient(conn), nil
}

func NewClientBlackWidow(config *config.TCPSocket) (natasha.BlackWidowClient, error) {
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

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.GetAddress(), config.GetPort()), options...)
	if err != nil {
		return nil, err
	}

	return natasha.NewBlackWidowClient(conn), nil
}
