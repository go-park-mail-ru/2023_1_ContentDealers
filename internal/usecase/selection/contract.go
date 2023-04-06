package selection

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context, limit, offset uint) ([]domain.Selection, error)
	GetByID(ctx context.Context, id uint64) (domain.Selection, error)
}

type ContentRepository interface {
	GetBySelectionIDs(ctx context.Context, IDs []uint64) (map[uint64]domain.Content, error)
}
