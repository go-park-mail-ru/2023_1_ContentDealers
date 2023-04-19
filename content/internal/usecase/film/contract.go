package film

import (
	"context"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type Repository interface {
	GetByContentID(ctx context.Context, id uint64) (domain2.Film, error)
}

type ContentUseCase interface {
	GetByID(ctx context.Context, id uint64) (domain2.Content, error)
}
