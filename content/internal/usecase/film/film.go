package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	repo    Repository
	content ContentUseCase
	logger  logging.Logger
}

func NewUseCase(repo Repository, content ContentUseCase, logger logging.Logger) *UseCase {
	return &UseCase{repo: repo, content: content, logger: logger}
}

func (uc *UseCase) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	film, err := uc.repo.GetByContentID(ctx, ContentID)
	if err != nil {
		return domain.Film{}, err
	}
	film.Content, err = uc.content.GetByID(ctx, ContentID)
	return film, err
}
