package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Genre, error)
	GetByID(ctx context.Context, id uint64) (domain.Genre, error)
}

type ContentRepository interface {
	GetByGenreOptions(ctx context.Context, options domain.ContentFilter) ([]domain.Content, error)
}
