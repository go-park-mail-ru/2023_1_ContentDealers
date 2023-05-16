package content

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	content.UnimplementedContentServiceServer

	useCase UseCase
	logger  logging.Logger
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{useCase: useCase, logger: logger}
}

func (service *Grpc) GetFilmByContentID(ctx context.Context, contentID *content.ContentID) (*content.Film, error) {
	id := contentID.GetID()

	foundFilm, err := service.useCase.GetFilmByContentID(ctx, id)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response content.Film
	err = dto.Map(&response, foundFilm)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetSeriesByContentID(ctx context.Context, contentID *content.ContentID) (*content.Series, error) {
	id := contentID.GetID()

	foundSeries, err := service.useCase.GetSeriesByContentID(ctx, id)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response content.Series
	err = dto.Map(&response, foundSeries)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
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
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response content.ContentSeq
	err = dto.Map(&response.Content, foundContent)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) AddRating(ctx context.Context,
	rating *content.Rating) (*content.Nothing, error) {
	err := service.useCase.AddRating(ctx, rating.ContentID, rating.Rating)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}
	return &content.Nothing{}, nil
}

func (service *Grpc) DeleteRating(ctx context.Context,
	rating *content.Rating) (*content.Nothing, error) {
	err := service.useCase.DeleteRating(ctx, rating.ContentID, rating.Rating)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}
	return &content.Nothing{}, nil
}
