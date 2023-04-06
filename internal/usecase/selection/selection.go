package selection

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Selection struct {
	repo    Repository
	content ContentRepository
}

func NewMovieSelection(repo Repository) *Selection {
	return &Selection{repo: repo}
}

func (uc *Selection) joinContent(ctx context.Context, selections *[]domain.Selection) error {

	IDs := make([]uint64, 0, len(*selections))
	IDToIdx := make(map[uint64]int, len(*selections))

	for idx, selection := range *selections {
		IDs = append(IDs, selection.ID)
		IDToIdx[selection.ID] = idx
	}

	SelectionIDContent, err := uc.content.GetBySelectionIDs(ctx, IDs)
	if err != nil {
		return err
	}

	for SelectionID, content := range SelectionIDContent {
		idx := IDToIdx[SelectionID]
		(*selections)[idx].Content = append((*selections)[idx].Content, content)
	}
	return nil
}

func (uc *Selection) GetAll(ctx context.Context, limit, offset uint) ([]domain.Selection, error) {
	selections, err := uc.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	err = uc.joinContent(ctx, &selections)

	return selections, err
}

func (uc *Selection) GetByID(ctx context.Context, id uint64) (domain.Selection, error) {
	selection, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Selection{}, err
	}

	selections := []domain.Selection{selection}
	err = uc.joinContent(ctx, &selections)

	return selections[0], err
}
