package favorites

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type ContentUseCase interface {
	GetByID(ctx context.Context, id uint64) (domain.Content, error)
}

type SessionUseCase interface {
	Create(ctx context.Context, user domain.User) (domain.Session, error)
	Get(ctx context.Context, sessionID string) (domain.Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type Gateway interface {
	Delete(ctx context.Context, favorite domain.FavoriteContent) error
	Add(ctx context.Context, favorite domain.FavoriteContent) error
	Get(ctx context.Context, options domain.FavoritesOptions) ([]domain.FavoriteContent, error)
}
