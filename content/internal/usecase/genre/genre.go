package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	repo    Repository
	content ContentRepository
	logger  logging.Logger
}

func NewUseCase(repo Repository, content ContentRepository, logger logging.Logger) *UseCase {
	return &UseCase{repo: repo, content: content, logger: logger}
}

// GetAll(ctx context.Context) ([]domain.Genre, error)
//	GetContentByOptions(ctx context.Context, options ContentFilter) ([]domain.Content, error)

func (uc *UseCase) GetAll(ctx context.Context) ([]domain.Genre, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UseCase) GetContentByOptions(ctx context.Context, options domain.ContentFilter) ([]domain.Content, error) {
	return uc.content.GetByGenreOptions(ctx, options)
}
