package selection

import (
	"context"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context, limit, offset uint) ([]domain2.Selection, error)
	GetByID(ctx context.Context, id uint64) (domain2.Selection, error)
}

type ContentRepository interface {
	GetBySelectionIDs(ctx context.Context, IDs []uint64) (map[uint64][]domain2.Content, error)
}
