package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type UseCase interface {
	GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error)
}
