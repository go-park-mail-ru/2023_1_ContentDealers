package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (domain.Content, error)
}

type PersonRolesRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain.PersonRoles, error)
}

type GenreRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Genre, error)
}

type SelectionRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Selection, error)
}

type CountryRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Country, error)
}
