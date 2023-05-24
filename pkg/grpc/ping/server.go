package ping

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/proto/ping"
)

type Grpc struct {
	ping.UnimplementedPingServiceServer
}

func NewGrpc() *Grpc {
	return &Grpc{}
}

func (service *Grpc) Ping(context.Context, *ping.Nothing) (*ping.Nothing, error) {
	return &ping.Nothing{}, nil
}
