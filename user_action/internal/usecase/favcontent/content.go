package favcontent

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type UseCase struct {
	repo   Repository
	logger logging.Logger
}

func NewUseCase(repo Repository, logger logging.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

func (uc *UseCase) Delete(ctx context.Context, favorite domain.FavoriteContent) error {
	return uc.repo.Delete(ctx, favorite)
}

func (uc *UseCase) Add(ctx context.Context, favorite domain.FavoriteContent) error {
	return uc.repo.Add(ctx, favorite)
}

func (uc *UseCase) HasFav(ctx context.Context, favorite domain.FavoriteContent) (bool, error) {
	return uc.repo.HasFav(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error) {
	return uc.repo.Get(ctx, options)
}
