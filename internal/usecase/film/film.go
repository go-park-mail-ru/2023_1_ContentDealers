package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Film struct {
	repo    Repository
	content ContentUseCase
}

func NewFilm(repo Repository, content ContentUseCase) *Film {
	return &Film{repo: repo, content: content}
}

func (uc *Film) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	film, err := uc.repo.GetByContentID(ctx, ContentID)
	if err != nil {
		return domain.Film{}, err
	}
	film.Content, err = uc.content.GetByID(ctx, ContentID)
	return film, err
}
