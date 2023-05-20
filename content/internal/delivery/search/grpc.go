package search

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	search.UnimplementedSearchServiceServer

	useCase UseCase
	logger  logging.Logger
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{useCase: useCase, logger: logger}
}

func (service *Grpc) Search(ctx context.Context, params *search.SearchParams) (*search.SearchResponse, error) {
	searchResult, err := service.useCase.Search(ctx, domain.SearchQuery{
		Query:  params.Query,
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response search.SearchResponse
	err = dto.Map(&response, searchResult)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}
