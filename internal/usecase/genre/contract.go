package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type ContentGateway interface {
	GetContentByOptions(ctx context.Context, filter domain.ContentFilter) ([]domain.Content, error)
	GetAllGenres(ctx context.Context) ([]domain.Genre, error)
}
