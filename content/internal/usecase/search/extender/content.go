package extender

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type ContentExtender struct {
	repo ContentRepository
}

func NewContentExtender(repo ContentRepository) *ContentExtender {
	return &ContentExtender{repo: repo}
}

func (extender *ContentExtender) Extend(ctx context.Context,
	query domain.SearchQuery) (func(search *domain.SearchResult), error) {
	content, err := extender.repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return func(search *domain.SearchResult) {
		search.SearchContent = content
	}, nil
}

func (extender *ContentExtender) GetSlug() string {
	return "content"
}
