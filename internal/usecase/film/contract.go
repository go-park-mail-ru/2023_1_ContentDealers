package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type ContentGateway interface {
	GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error)
}