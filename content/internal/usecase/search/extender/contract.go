package extender

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type ContentRepository interface {
	Search(ctx context.Context, query domain.SearchQuery) (domain.SearchContent, error)
}

type PersonRepository interface {
	Search(ctx context.Context, query domain.SearchQuery) (domain.SearchPerson, error)
}
