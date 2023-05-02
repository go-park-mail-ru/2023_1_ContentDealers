package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase interface {
	GetGenreContent(ctx context.Context, filter domain.ContentFilter) (domain.GenreContent, error)
	GetAll(ctx context.Context) ([]domain.Genre, error)
}
