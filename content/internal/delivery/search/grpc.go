package search

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/search"
)

type Grpc struct {
	search.UnimplementedSearchServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

func (service *Grpc) Search(ctx context.Context, params *search.SearchParams) (*search.SearchResponse, error) {
	searchResult, err := service.useCase.Search(ctx, params.GetQuery())
	if err != nil {
		return nil, err
	}

	var response search.SearchResponse
	err = dto.Map(&response, searchResult)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
