package favorites

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	domainRate "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type SessionGateway interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type Gateway interface {
	DeleteRating(ctx context.Context, rating domainRate.Rating) (domainRate.Rating, error)
	AddRating(ctx context.Context, rating domainRate.Rating) error
	HasRating(ctx context.Context, rating domainRate.Rating) (domainRate.HasRating, error)
	GetRatingByUser(ctx context.Context, options domainRate.RatingsOptions) (domainRate.Ratings, error)
	GetRatingByContent(ctx context.Context, options domainRate.RatingsOptions) (domainRate.Ratings, error)
}

type ContentGateway interface {
	GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domain.Content, error)
	AddRating(ctx context.Context, ContentID uint64, rating float32) error
	DeleteRating(ctx context.Context, ContentID uint64, rating float32) error
}
