package film

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/genre"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	genre.UnimplementedGenreServiceServer

	useCase UseCase
	logger  logging.Logger
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{useCase: useCase, logger: logger}
}

func (service *Grpc) GetContentByOptions(ctx context.Context, options *genre.Options) (*genre.GenreContent, error) {
	genreContent, err := service.useCase.GetGenreContent(ctx, domain.ContentFilter{
		ID:     options.GetID(),
		Limit:  options.GetLimit(),
		Offset: options.GetOffset(),
	})
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response genre.GenreContent
	err = dto.Map(&response, genreContent)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetAllGenres(ctx context.Context, _ *genre.Nothing) (*genre.Genres, error) {
	genres, err := service.useCase.GetAll(ctx)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response genre.Genres
	err = dto.Map(&response.Genres, genres)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}
