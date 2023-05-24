package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (domain.Content, error)
	GetFilmByContentID(ctx context.Context, id uint64) (domain.Film, error)
	GetSeriesByContentID(ctx context.Context, id uint64) (domain.Series, error)
	GetByIDs(ctx context.Context, ids []uint64) ([]domain.Content, error)
	AddRating(ctx context.Context, ContentID uint64, rating float32) error
	DeleteRating(ctx context.Context, ContentID uint64, rating float32) error
	GetEpisodesBySeasonNum(ctx context.Context, ContentID uint64, seasonNum uint32) ([]domain.Episode, error)
}

type PersonRolesUseCase interface {
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
