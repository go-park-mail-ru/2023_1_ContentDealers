package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Person struct {
	repo Repository
}

func NewPerson(repo Repository) *Person {
	return &Person{repo: repo}
}

func (uc *Person) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	return uc.repo.GetByID(ctx, id)
}
