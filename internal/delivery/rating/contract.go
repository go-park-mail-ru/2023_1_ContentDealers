package rating

import (
	"context"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type RatingUseCase interface {
	Delete(ctx context.Context, rating domain.Rating) error
	Add(ctx context.Context, rating domain.Rating) error
	Has(ctx context.Context, rating domain.Rating) (domain.HasRating, error)
	GetByUser(ctx context.Context, options domain.RatingsOptions) ([]domainContent.Content, []domain.Rating, bool, error)
	GetByContent(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error)
}
