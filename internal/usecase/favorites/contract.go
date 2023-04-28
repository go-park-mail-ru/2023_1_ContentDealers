package favorites

import (
	"context"

	domainFav "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type ContentUseCase interface {
	GetByID(ctx context.Context, id uint64) (domain.Content, error)
}

type SessionUseCase interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type Gateway interface {
	Delete(ctx context.Context, favorite domainFav.FavoriteContent) error
	Add(ctx context.Context, favorite domainFav.FavoriteContent) error
	Get(ctx context.Context, options domainFav.FavoritesOptions) ([]domainFav.FavoriteContent, error)
}
