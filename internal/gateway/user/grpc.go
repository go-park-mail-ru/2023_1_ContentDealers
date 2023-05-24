package user

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/dranikpg/dto-mapper"
	interceptorClient "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/client"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	userProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	dirAvatars             = "media/avatars"
	allPerms   os.FileMode = 0777
	bufSize                = 1024
)

type Gateway struct {
	logger      logging.Logger
	userManager userProto.UserServiceClient
}

func NewGateway(logger logging.Logger, cfg ServiceUserConfig) (*Gateway, error) {
	interceptor := interceptorClient.NewInterceptorClient("user", logger)

	grpcConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
		return nil, err
	}

	userManager := userProto.NewUserServiceClient(grpcConn)

	err = ping.Ping(grpcConn)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &Gateway{logger: logger, userManager: userManager}, nil
}

func (gate *Gateway) Register(ctx context.Context, user domain.User) (domain.User, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	userRequest := userProto.User{}
	dto.Map(&userRequest, user)
	userResponse, err := gate.userManager.Register(ctx, &userRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		switch {
		case strings.Contains(err.Error(), domain.ErrNotValidPassword.Error()):
			return domain.User{}, domain.ErrNotValidPassword
		case strings.Contains(err.Error(), domain.ErrNotValidEmail.Error()):
			return domain.User{}, domain.ErrNotValidEmail
		case strings.Contains(err.Error(), domain.ErrUserAlreadyExists.Error()):
			return domain.User{}, domain.ErrUserAlreadyExists
		default:
			return domain.User{}, err
		}
	}
	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}
	user.SubscriptionExpiryDate = userResponse.SubscriptionExpiryDate.AsTime()
	return user, nil
}

func (gate *Gateway) Auth(ctx context.Context, user domain.User) (domain.User, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	userRequest := userProto.User{}
	dto.Map(&userRequest, user)
	userResponse, err := gate.userManager.Auth(ctx, &userRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		switch {
		case strings.Contains(err.Error(), domain.ErrWrongCredentials.Error()):
			return domain.User{}, domain.ErrWrongCredentials
		default:
			return domain.User{}, err
		}
	}
	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}
	user.SubscriptionExpiryDate = userResponse.SubscriptionExpiryDate.AsTime()
	return user, nil
}

func (gate *Gateway) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	var UserIDRequest userProto.ID
	UserIDRequest.ID = id

	userResponse, err := gate.userManager.GetByID(ctx, &UserIDRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		switch {
		case strings.Contains(err.Error(), domain.ErrUserNotFound.Error()):
			return domain.User{}, domain.ErrUserNotFound
		default:
			return domain.User{}, err
		}
	}

	user := domain.User{}

	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}

	user.SubscriptionExpiryDate = userResponse.SubscriptionExpiryDate.AsTime()
	return user, nil
}

func (gate *Gateway) Update(ctx context.Context, user domain.User) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	userRequest := userProto.User{}
	dto.Map(&userRequest, user)
	_, err := gate.userManager.Update(ctx, &userRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		switch {
		case strings.Contains(err.Error(), domain.ErrUserAlreadyExists.Error()):
			return domain.ErrUserAlreadyExists
		default:
			return err
		}
	}
	return nil
}

func (gate *Gateway) UpdateAvatar(ctx context.Context, user domain.User, reader io.Reader) (domain.User, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := gate.userManager.UpdateAvatar(ctx)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}

	userRequest := userProto.User{}
	dto.Map(&userRequest, user)

	err = stream.Send(&userProto.UserAvatar{
		User: &userRequest,
	})
	if err != nil {
		gate.logger.WithRequestID(ctx).Tracef("Error sending user object: %v", err)
		return domain.User{}, err
	}

	buffer := make([]byte, bufSize)

	for {
		// nolint:govet
		bytesRead, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			gate.logger.WithRequestID(ctx).Trace(err)
			return domain.User{}, err
		}

		err = stream.Send(&userProto.UserAvatar{
			Chunk: buffer[:bytesRead],
		})
		if err != nil {
			gate.logger.WithRequestID(ctx).Tracef("Error sending chunk to server: %v", err)
			return domain.User{}, err
		}
	}

	userResponse, err := stream.CloseAndRecv()
	if err != nil {
		gate.logger.WithRequestID(ctx).Tracef("Error receiving response from server: %v", err)
		return domain.User{}, err
	}

	dto.Map(&user, userResponse)
	return user, nil
}

func (gate *Gateway) DeleteAvatar(ctx context.Context, user domain.User) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	userRequest := userProto.User{}
	dto.Map(&userRequest, user)
	_, err := gate.userManager.DeleteAvatar(ctx, &userRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}
