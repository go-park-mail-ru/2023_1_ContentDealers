package favcontent

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

const (
	minRating = 1
	maxRating = 10
)

type UseCase struct {
	repo   Repository
	logger logging.Logger
}

func NewUseCase(repo Repository, logger logging.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

func (uc *UseCase) Delete(ctx context.Context, rating domain.Rating) (domain.Rating, error) {
	return uc.repo.Delete(ctx, rating)
}

func (uc *UseCase) Add(ctx context.Context, rating domain.Rating) error {
	if rating.Rating < minRating || rating.Rating > maxRating {
		return fmt.Errorf("rating out of range")
	}
	return uc.repo.Add(ctx, rating)
}

func (uc *UseCase) Has(ctx context.Context, rating domain.Rating) (domain.HasRating, error) {
	return uc.repo.Has(ctx, rating)
}

func (uc *UseCase) GetByUser(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	return uc.repo.GetByUser(ctx, options)
}

func (uc *UseCase) GetByContent(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	return uc.repo.GetByUser(ctx, options)
}
