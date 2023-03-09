package usecase

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

var _ contract.MovieSelectionUseCase = (*MovieSelection)(nil)

type MovieSelection struct {
	repo contract.MovieSelectionRepository
}

func NewMovieSelection(repo contract.MovieSelectionRepository) *MovieSelection {
	return &MovieSelection{repo: repo}
}

func (uc *MovieSelection) GetAll() ([]domain.MovieSelection, error) {
	return uc.repo.GetAll()
}

func (uc *MovieSelection) GetByID(id uint64) (domain.MovieSelection, error) {
	return uc.repo.GetByID(id)
}
