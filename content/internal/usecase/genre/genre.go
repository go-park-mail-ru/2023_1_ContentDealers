package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase struct {
	repo    Repository
	content ContentRepository
}

func NewUseCase(repo Repository, content ContentRepository) *UseCase {
	return &UseCase{repo: repo, content: content}
}

// GetAll(ctx context.Context) ([]domain.Genre, error)
//	GetContentByOptions(ctx context.Context, options ContentFilter) ([]domain.Content, error)

func (uc *UseCase) GetAll(ctx context.Context) ([]domain.Genre, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UseCase) GetGenreContent(ctx context.Context, options domain.ContentFilter) (domain.GenreContent, error) {
	result := domain.GenreContent{}
	var err error
	result.Content, err = uc.content.GetByGenreOptions(ctx, options)
	if err != nil {
		return result, err
	}
	result.Genre, err = uc.repo.GetByID(ctx, options.ID)
	return result, err
}
