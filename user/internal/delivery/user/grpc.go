package user

import (
	"bytes"
	"context"
	"io"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/internal/domain"
	userProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, err
	}
	newUser, err := service.userUseCase.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func (service *Grpc) Auth(ctx context.Context, userRequest *userProto.User) (*userProto.User, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		return nil, err
	}
	newUser, err := service.userUseCase.Auth(ctx, user)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func (service *Grpc) GetByUserID(ctx context.Context, IDRequest *userProto.ID) (*userProto.User, error) {
	newUser, err := service.userUseCase.GetByID(ctx, IDRequest.ID)
	if err != nil {
		return nil, err
	}
	userResponse := userProto.User{}
	err = dto.Map(&userResponse, newUser)
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func (service *Grpc) Update(ctx context.Context, userRequest *userProto.User) (*userProto.Nothing, error) {
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
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

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "Error reading the client stream: %v", err)
		}

		if chunkCount == 0 {
			userRequest = chunk.GetUser()
		}

		chunkCount++
		chunkData := chunk.GetChunk()

		bytesAvatar = append(bytesAvatar, chunkData...)
		if err != nil {
			return status.Errorf(codes.Internal, "Could not store the image: %v", err)
		}
	}
	reader := bytes.NewReader(bytesAvatar)
	user := domain.User{}
	err := dto.Map(&user, userRequest)
	if err != nil {
		return err
	}

	user, err = service.userUseCase.UpdateAvatar(ctx, user, reader)
	userResponse := &userProto.User{}
	err = dto.Map(userResponse, user)
	if err != nil {
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

func (service *Grpc) Ping(ctx context.Context, req *userProto.PingRequest) (*userProto.PingResponse, error) {
	service.logger.Info("Ping reached session microservice")
	return &userProto.PingResponse{}, nil
}
