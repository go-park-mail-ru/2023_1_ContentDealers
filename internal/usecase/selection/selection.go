package selection

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type MovieSelection struct {
	repo MovieSelectionRepository
}

func NewMovieSelection(repo MovieSelectionRepository) *MovieSelection {
	return &MovieSelection{repo: repo}
}

func (uc *MovieSelection) GetAll() ([]domain.Selection, error) {
	return uc.repo.GetAll()
}

func (uc *MovieSelection) GetByID(id uint64) (domain.Selection, error) {
	return uc.repo.GetByID(id)
}