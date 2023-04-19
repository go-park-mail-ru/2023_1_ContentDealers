package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Film struct {
	repo    Repository
	content ContentUseCase
	logger  logging.Logger
}

func NewFilm(repo Repository, content ContentUseCase, logger logging.Logger) *Film {
	return &Film{repo: repo, content: content, logger: logger}
}

func (uc *Film) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	film, err := uc.repo.GetByContentID(ctx, ContentID)
	if err != nil {
		return domain.Film{}, err
	}
	film.Content, err = uc.content.GetByID(ctx, ContentID)
	return film, err
}
