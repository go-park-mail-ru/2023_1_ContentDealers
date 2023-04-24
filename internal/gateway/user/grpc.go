package user

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	userProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	dirAvatars             = "media/avatars"
	allPerms   os.FileMode = 0777
)

type Gateway struct {
	logger      logging.Logger
	userManager userProto.UserServiceClient
	interseptor UserInterceptor
}

func pingServer(ctx context.Context, client userProto.UserServiceClient) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := client.Ping(ctx, &userProto.PingRequest{})
	if err != nil {
		return err
	}

	return nil
}

func NewGateway(logger logging.Logger, cfg ServiceUserConfig) (*Gateway, error) {
	interseptor := UserInterceptor{logger: logger}

	grcpConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interseptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
		return nil, err
	}

	userManager := userProto.NewUserServiceClient(grcpConn)

	err = pingServer(context.Background(), userManager)
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
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}
	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}
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
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}
	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}
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
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}

	user := domain.User{}

	err = dto.Map(&user, userResponse)
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}
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
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return err
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
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.User{}, err
	}

	userRequest := userProto.User{}
	dto.Map(&userRequest, user)

	err = stream.Send(&userProto.UserAvatar{
		User: &userRequest,
	})
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Tracef("Error sending user object: %w", err)
		return domain.User{}, err
	}

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			gate.logger.WithFields(logrus.Fields{
				"request_id": ctx.Value("requestID").(string),
			}).Trace(err)
			return domain.User{}, err
		}

		err = stream.Send(&userProto.UserAvatar{
			Chunk: buffer[:bytesRead],
		})
		if err != nil {
			gate.logger.WithFields(logrus.Fields{
				"request_id": ctx.Value("requestID").(string),
			}).Tracef("Error sending chunk to server: %w", err)
			return domain.User{}, err
		}
	}

	userResponse, err := stream.CloseAndRecv()
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Tracef("Error receiving response from server: %w", err)
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
	_, err := gate.userManager.Update(ctx, &userRequest)
	if err != nil {
		gate.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return err
	}
	return nil
}
