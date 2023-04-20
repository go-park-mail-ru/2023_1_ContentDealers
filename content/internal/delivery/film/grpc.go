package film

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/film"
)

type Grpc struct {
	film.UnimplementedFilmServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

func (service *Grpc) GetByContentID(ctx context.Context, contentID *film.ContentID) (*film.Film, error) {
	id := contentID.GetID()

	foundFilm, err := service.useCase.GetByContentID(ctx, id)
	if err != nil {
		return nil, err
	}

	var response film.Film
	err = dto.Map(&response, foundFilm)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
