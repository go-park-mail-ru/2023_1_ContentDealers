package ping

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/proto/ping"
	"google.golang.org/grpc"
)

const defaultTimeout = time.Second * 5

func Ping(grpcConn *grpc.ClientConn) error {
	return PingWithTimeout(grpcConn, defaultTimeout)
}

func PingWithTimeout(grpcConn *grpc.ClientConn, timeout time.Duration) error {
	pingService := ping.NewPingServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := pingService.Ping(ctx, &ping.Nothing{})
	return err
}
