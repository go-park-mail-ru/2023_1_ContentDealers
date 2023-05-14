package history_views

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

func (uc *UseCase) UpdateProgressView(ctx context.Context, view domain.View) error {
	return uc.repo.UpdateProgressView(ctx, view)
}

func (uc *UseCase) GetPartiallyViewsByUser(ctx context.Context, viewOptions domain.ViewsOptions) (domain.Views, error) {
	return uc.repo.GetPartiallyViewsByUser(ctx, viewOptions)
}

func (uc *UseCase) GetAllViewsByUser(ctx context.Context, viewOptions domain.ViewsOptions) (domain.Views, error) {
	return uc.repo.GetAllViewsByUser(ctx, viewOptions)
}

func (uc *UseCase) HasView(ctx context.Context, view domain.View) (domain.HasView, error) {
	return uc.repo.HasView(ctx, view)
}
