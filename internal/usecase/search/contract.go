package search

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type ContentGateway interface {
	Search(ctx context.Context, query string) (domain.Search, error)
}
