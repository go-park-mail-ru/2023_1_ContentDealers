package user

import (
	"context"

	interceptorClient "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/client"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
	"google.golang.org/grpc"
)

type Gateway struct {
	userService user.UserServiceClient
	logger      logging.Logger
}

func NewGateway(cfg ServiceUserConfig, logger logging.Logger) (*Gateway, error) {
	interceptor := interceptorClient.NewInterceptorClient("content", logger)

	grpcConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.AccessLog),
	)
	if err != nil {
		return &Gateway{}, err
	}

	result := Gateway{}
	result.userService = user.NewUserServiceClient(grpcConn)
	result.logger = logger

	err = ping.Ping(grpcConn)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (service *Gateway) Subscribe(ctx context.Context, userID uint64) error {
	_, err := service.userService.Subscribe(ctx, &user.User{ID: userID})
	return err
}
