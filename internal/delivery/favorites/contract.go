package favorites

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type FavContentUseCase interface {
	Delete(ctx context.Context, favorite domain.FavoriteContent) error
	Add(ctx context.Context, favorite domain.FavoriteContent) error
	Get(ctx context.Context, options domain.FavoritesOptions) ([]domain.FavoriteContent, error)
}
