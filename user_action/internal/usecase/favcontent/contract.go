package favcontent

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type Repository interface {
	Delete(ctx context.Context, favorite domain.FavoriteContent) error
	Add(ctx context.Context, favorite domain.FavoriteContent) error
	HasFav(ctx context.Context, favorite domain.FavoriteContent) (bool, error)
	Get(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error)
}
