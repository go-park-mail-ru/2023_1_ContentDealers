package favorites

import (
	"context"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
)

type FavContentUseCase interface {
	Delete(ctx context.Context, favorite domain.FavoriteContent) error
	Add(ctx context.Context, favorite domain.FavoriteContent) error
	Get(ctx context.Context, options domain.FavoritesOptions) ([]domainContent.Content, error)
}
