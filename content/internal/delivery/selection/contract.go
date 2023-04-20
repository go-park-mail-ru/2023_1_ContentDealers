package selection

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type UseCase interface {
	GetAll(ctx context.Context, limit, offset uint) ([]domain.Selection, error)
	GetByID(ctx context.Context, id uint64) (domain.Selection, error)
}
