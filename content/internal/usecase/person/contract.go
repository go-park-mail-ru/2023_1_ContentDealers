package person

import (
	"context"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (domain2.Person, error)
}

type ContentRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain2.Content, error)
}

type RoleRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain2.Role, error)
}

type GenreRepository interface {
	GetByPersonID(ctx context.Context, PersonID uint64) ([]domain2.Genre, error)
}
