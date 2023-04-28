package extender

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type ContentExtender struct {
	repo   ContentRepository
	logger logging.Logger
}

func NewContentExtender(repo ContentRepository, logger logging.Logger) *ContentExtender {
	return &ContentExtender{repo: repo, logger: logger}
}

func (extender *ContentExtender) Extend(ctx context.Context, query string) func(search *domain.Search) error {
	return func(search *domain.Search) error {
		content, err := extender.repo.Search(ctx, query)
		if err != nil {
			return err
		}
		search.Content = content
		return nil
	}
}
