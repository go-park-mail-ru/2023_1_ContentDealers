package usecase

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type MovieSelectionRepository interface {
	GetAll() ([]domain.MovieSelection, error)
	GetById(id uint64) (domain.MovieSelection, error)
}

type MovieSelectionUseCase struct {
	repo MovieSelectionRepository
}

func NewMovieSelectionUseCase(repo MovieSelectionRepository) MovieSelectionUseCase {
	return MovieSelectionUseCase{repo: repo}
}

func (uc *MovieSelectionUseCase) GetAll() ([]domain.MovieSelection, error) {
	return uc.repo.GetAll()
}

func (uc *MovieSelectionUseCase) GetById(id uint64) (domain.MovieSelection, error) {
	return uc.repo.GetById(id)
}
