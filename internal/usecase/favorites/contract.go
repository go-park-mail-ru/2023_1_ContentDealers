package favorites

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	domainFav "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type SessionGateway interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type Gateway interface {
	DeleteFavContent(ctx context.Context, favorite domainFav.FavoriteContent) error
	AddFavContent(ctx context.Context, favorite domainFav.FavoriteContent) error
	GetFavContent(ctx context.Context, options domainFav.FavoritesOptions) (domainFav.FavoritesContent, error)
	HasFavContent(ctx context.Context, favorite domainFav.FavoriteContent) (bool, error)
}

type ContentGateway interface {
	GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domain.Content, error)
}
