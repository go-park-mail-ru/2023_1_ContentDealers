package search

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Search struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewSearch(contentGateway ContentGateway, logger logging.Logger) *Search {
	return &Search{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *Search) Search(ctx context.Context, query string) (domain.Search, error) {
	return uc.contentGateway.Search(ctx, query)
}
