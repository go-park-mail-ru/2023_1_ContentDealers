package user

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	userProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Grpc struct {
	userProto.UnimplementedUserServiceServer
	userUseCase UserUseCase
	logger      logging.Logger
}

func NewGrpc(userUseCase UserUseCase, logger logging.Logger) *Grpc {
	return &Grpc{userUseCase: userUseCase, logger: logger}
}

func (service *Grpc) Register(ctx context.Context, userRequest *userProto.User) (*userProto.User, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	newUser, err := service.userUseCase.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	userResponse.SubscriptionExpiryDate = timestamppb.New(newUser.SubscriptionExpiryDate)
	return &userResponse, nil
}

func (service *Grpc) Auth(ctx context.Context, userRequest *userProto.User) (*userProto.User, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	newUser, err := service.userUseCase.Auth(ctx, user)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	userResponse.SubscriptionExpiryDate = timestamppb.New(newUser.SubscriptionExpiryDate)
	return &userResponse, nil
}

func (service *Grpc) GetByID(ctx context.Context, IDRequest *userProto.ID) (*userProto.User, error) {
	newUser, err := service.userUseCase.GetByID(ctx, IDRequest.ID)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	userResponse.SubscriptionExpiryDate = timestamppb.New(newUser.SubscriptionExpiryDate)
	return &userResponse, nil
}

func (service *Grpc) Update(ctx context.Context, userRequest *userProto.User) (*userProto.Nothing, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	err = service.userUseCase.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userProto.Nothing{}, nil
}

func (service *Grpc) UpdateAvatar(stream userProto.UserService_UpdateAvatarServer) error {
	userRequest := &userProto.User{}
	chunkCount := 0
	var bytesAvatar []byte
	ctx := stream.Context()

	var reqID string
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists || len(md.Get("requestId")) == 0 {
		reqID = "unknown"
	} else {
		reqID = md.Get("requestId")[0]
	}
	ctx = context.WithValue(ctx, "requestID", reqID)

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			err := fmt.Errorf("Error reading the client stream: %w", err)
			service.logger.WithRequestID(ctx).Trace(err)
			return status.Error(codes.Unknown, err.Error())
		}

		if chunkCount == 0 {
			userRequest = chunk.GetUser()
		}

		chunkCount++
		chunkData := chunk.GetChunk()

		bytesAvatar = append(bytesAvatar, chunkData...)
	}
	reader := bytes.NewReader(bytesAvatar)
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return err
	}

	user, err = service.userUseCase.UpdateAvatar(ctx, user, reader)
	userResponse := &userProto.User{}
	err = dto.Map(userResponse, user)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	stream.SendAndClose(userResponse)
	return nil
}

func (service *Grpc) DeleteAvatar(ctx context.Context, userRequest *userProto.User) (*userProto.Nothing, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		return nil, err
	}
	err = service.userUseCase.DeleteAvatar(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userProto.Nothing{}, nil
}

func (service *Grpc) Subscribe(ctx context.Context, userRequest *userProto.User) (*userProto.Nothing, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		return nil, err
	}
	err = service.userUseCase.Subscribe(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userProto.Nothing{}, nil
}
