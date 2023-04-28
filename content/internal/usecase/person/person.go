package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	repo    Repository
	content ContentRepository
	role    RoleRepository
	genre   GenreRepository
	logger  logging.Logger
}

type Options struct {
	Repo    Repository
	Content ContentRepository
	Role    RoleRepository
	Genre   GenreRepository
	Logger  logging.Logger
}

func NewUseCase(options Options) *UseCase {
	return &UseCase{
		repo:    options.Repo,
		content: options.Content,
		role:    options.Role,
		genre:   options.Genre,
		logger:  options.Logger,
	}
}

func (uc *UseCase) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
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
