package film

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/genre"
)

type Grpc struct {
	genre.UnimplementedGenreServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

func (service *Grpc) GetContentByOptions(ctx context.Context, options *genre.Options) (*genre.GenreContent, error) {
	content, err := service.useCase.GetContentByOptions(ctx, domain.ContentFilter{
		ID:     options.GetID(),
		Limit:  options.GetLimit(),
		Offset: options.GetOffset(),
	})
	if err != nil {
		return nil, err
	}

	var response genre.GenreContent
	err = dto.Map(&response, content)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetAllGenres(ctx context.Context, _ *genre.Nothing) (*genre.Genres, error) {
	genres, err := service.useCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var response genre.Genres
	err = dto.Map(&response.Genres, genres)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
