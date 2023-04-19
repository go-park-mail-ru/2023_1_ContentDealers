package content

import (
	"context"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (domain2.Content, error)
}

type PersonRolesUseCase interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain2.PersonRoles, error)
}

type GenreRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain2.Genre, error)
}

type SelectionRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain2.Selection, error)
}

type CountryRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain2.Country, error)
}
