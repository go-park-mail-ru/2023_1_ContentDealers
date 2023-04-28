package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
)

type Repository interface {
	Delete(ctx context.Context, favorite domain.FavoriteContent) error
	Add(ctx context.Context, favorite domain.FavoriteContent) error
	Get(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error)
}
