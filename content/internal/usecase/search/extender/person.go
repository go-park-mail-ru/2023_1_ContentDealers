package extender

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type PersonExtender struct {
	repo PersonRepository
}

func NewPersonExtender(repo PersonRepository) *PersonExtender {
	return &PersonExtender{repo: repo}
}

func (extender *PersonExtender) Extend(ctx context.Context,
	query domain.SearchQuery) (func(search *domain.SearchResult), error) {
	persons, err := extender.repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return func(search *domain.SearchResult) {
		search.SearchPerson = persons
	}, nil
}

func (extender *PersonExtender) GetSlug() string {
	return "person"
}
