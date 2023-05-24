package search

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type Extender interface {
	Extend(ctx context.Context, query domain.SearchQuery) (func(search *domain.SearchResult), error)
	GetSlug() string
}
