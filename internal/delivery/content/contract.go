package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase interface {
	GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error)
	GetSeriesByContentID(ctx context.Context, ContentID uint64) (domain.Series, error)
}
