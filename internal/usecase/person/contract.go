package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (domain.Person, error)
}

type ContentRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain.Content, error)
}

type RoleRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain.Role, error)
}

type GenreRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain.Genre, error)
}
