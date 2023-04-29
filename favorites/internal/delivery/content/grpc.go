package user

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
	favProto "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	favProto.UnimplementedFavoritesContentServiceServer
	favUseCase FavContentUseCase
	logger     logging.Logger
}

func NewGrpc(favUseCase FavContentUseCase, logger logging.Logger) *Grpc {
	return &Grpc{favUseCase: favUseCase, logger: logger}
}

func (service *Grpc) DeleteContent(ctx context.Context, favRequest *favProto.Favorite) (*favProto.Nothing, error) {
	fav := domain.FavoriteContent{}
	err := dto.Map(&fav, favRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	err = service.favUseCase.Delete(ctx, fav)
	if err != nil {
		return nil, err
	}
	return &favProto.Nothing{}, nil
}

func (service *Grpc) AddContent(ctx context.Context, favRequest *favProto.Favorite) (*favProto.Nothing, error) {
	fav := domain.FavoriteContent{}
	err := dto.Map(&fav, favRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	err = service.favUseCase.Add(ctx, fav)
	if err != nil {
		return nil, err
	}
	return &favProto.Nothing{}, nil
}

func (service *Grpc) HasFavContent(ctx context.Context, favRequest *favProto.Favorite) (*favProto.HasFav, error) {
	fav := domain.FavoriteContent{}
	err := dto.Map(&fav, favRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	HasFav, err := service.favUseCase.HasFav(ctx, fav)
	if err != nil {
		return nil, err
	}
	return &favProto.HasFav{HasFav: HasFav}, nil
}

func (service *Grpc) GetContent(ctx context.Context, favOptionsRequest *favProto.FavoritesOptions) (*favProto.Favorites, error) {
	favOptions := domain.FavoritesOptions{}
	err := dto.Map(&favOptions, favOptionsRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	favorites, err := service.favUseCase.Get(ctx, favOptions)
	if err != nil {
		return nil, err
	}
	favoritesResponse := favProto.Favorites{}
	err = dto.Map(&favoritesResponse, favorites)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	return &favoritesResponse, nil
}

func (service *Grpc) Ping(ctx context.Context, req *favProto.PingRequest) (*favProto.PingResponse, error) {
	service.logger.Info("Ping reached favorites microservice for content")
	return &favProto.PingResponse{}, nil
}
