package content

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/content"
)

type Grpc struct {
	content.UnimplementedContentServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

func (service *Grpc) GetFilmByContentID(ctx context.Context, contentID *content.ContentID) (*content.Film, error) {
	id := contentID.GetID()

	foundFilm, err := service.useCase.GetFilmByContentID(ctx, id)
	if err != nil {
		return nil, err
	}

	var response content.Film
	err = dto.Map(&response, foundFilm)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetSeriesByContentID(ctx context.Context, contentID *content.ContentID) (*content.Series, error) {
	id := contentID.GetID()

	foundSeries, err := service.useCase.GetSeriesByContentID(ctx, id)
	if err != nil {
		return nil, err
	}

	var response content.Series
	err = dto.Map(&response, foundSeries)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetContentByContentIDs(ctx context.Context,
	contentIDs *content.ContentIDs) (*content.ContentSeq, error) {
	var ids []uint64
	for _, id := range contentIDs.GetContentIDs() {
		ids = append(ids, id.GetID())
	}

	foundContent, err := service.useCase.GetContentByContentIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	var response content.ContentSeq
	err = dto.Map(&response.Content, foundContent)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
