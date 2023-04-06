package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Person struct {
	repo    Repository
	content ContentRepository
	role    RoleRepository
	genre   GenreRepository
}

func NewPerson(repo Repository) *Person {
	return &Person{repo: repo}
}

func (uc *Person) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	person, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Person{}, err
	}
	person.ParticipatedIn, err = uc.content.GetByPersonID(ctx, id)
	if err != nil {
		return domain.Person{}, err
	}
	person.Roles, err = uc.role.GetByPersonID(ctx, id)
	if err != nil {
		return domain.Person{}, err
	}
	person.Genres, err = uc.genre.GetByPersonID(ctx, id)
	if err != nil {
		return domain.Person{}, err
	}
	return person, nil
}
