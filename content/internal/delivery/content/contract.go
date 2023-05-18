package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase interface {
	GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error)
	GetSeriesByContentID(ctx context.Context, ContentID uint64) (domain.Series, error)
	GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domain.Content, error)
	AddRating(ctx context.Context, ContentID uint64, rating float32) error
	DeleteRating(ctx context.Context, ContentID uint64, rating float32) error
	GetEpisodesBySeasonNum(ctx context.Context, ContentID uint64, seasonNum uint32) ([]domain.Episode, error)
}
