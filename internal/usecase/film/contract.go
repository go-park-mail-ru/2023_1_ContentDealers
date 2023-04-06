package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository interface {
	GetByContentID(ctx context.Context, id uint64) (domain.Film, error)
}

type ContentUseCase interface {
	GetByID(ctx context.Context, id uint64) (domain.Content, error)
}
