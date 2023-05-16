package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type Options struct {
	ID     uint64
	Limit  uint32
	Offset uint32
}

type UseCase interface {
	GetAll(ctx context.Context) ([]domain.Genre, error)
	GetGenreContent(ctx context.Context, options domain.ContentFilter) (domain.GenreContent, error)
}
