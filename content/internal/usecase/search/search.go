package search

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"golang.org/x/sync/errgroup"
)

type Search struct {
	extenders []Extender
	logger    logging.Logger
}

func NewSearch(extenders []Extender, logger logging.Logger) *Search {
	return &Search{extenders: extenders, logger: logger}
}

func (uc *Search) Search(ctx context.Context, query string) (domain.Search, error) {
	var result domain.Search
	var mu sync.Mutex
	group, ctx := errgroup.WithContext(ctx)

	for _, extender := range uc.extenders {
		group.Go(func() error {
			applier := extender.Extend(ctx, query)

			mu.Lock()
			defer mu.Unlock()

			return applier(&result)
		})
	}

	err := group.Wait()
	if err != nil {
		return result, err
	}
	return result, nil
}
