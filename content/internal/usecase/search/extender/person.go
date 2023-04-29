package extender

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type PersonExtender struct {
	repo   PersonRepository
	logger logging.Logger
}

func NewPersonExtender(repo PersonRepository, logger logging.Logger) *PersonExtender {
	return &PersonExtender{repo: repo, logger: logger}
}

func (extender *PersonExtender) Extend(ctx context.Context, query string) (func(search *domain.Search), error) {
	persons, err := extender.repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return func(search *domain.Search) {
		search.Persons = persons
	}, nil
}
