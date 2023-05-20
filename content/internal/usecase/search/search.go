package search

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase struct {
	extenders []Extender
}

func NewUseCase(extenders []Extender) *UseCase {
	return &UseCase{extenders: extenders}
}

func (uc *UseCase) Search(ctx context.Context, query domain.SearchQuery) (domain.SearchResult, error) {
	var result domain.SearchResult
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	for _, extender := range uc.extenders {
		if query.TargetSlug == "" || query.TargetSlug == extender.GetSlug() {
			wg.Add(1)
			go func(extender Extender) {
				defer wg.Done()
				applier, err := extender.Extend(ctx, query)
				if err != nil {
					return
				}

				mu.Lock()
				defer mu.Unlock()
				applier(&result)
			}(extender)
		}

	}

	wg.Wait()
	return result, nil
}
