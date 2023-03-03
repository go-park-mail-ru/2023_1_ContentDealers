package repository

import (
	"errors"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

var ErrMovieSelectionNotFound = errors.New("movie selection not found")

var _ usecase.MovieSelectionRepository = (*MovieSelectionInMemoryRepository)(nil)

type MovieSelectionInMemoryRepository struct {
	storage []domain.MovieSelection
}

func NewMovieSelectionInMemoryRepository() MovieSelectionInMemoryRepository {
	return MovieSelectionInMemoryRepository{}
}

func (repo *MovieSelectionInMemoryRepository) Add(selections domain.MovieSelection) {
	repo.storage = append(repo.storage, selections)
}

func (repo *MovieSelectionInMemoryRepository) GetAll() ([]domain.MovieSelection, error) {
	return repo.storage, nil
}

func (repo *MovieSelectionInMemoryRepository) GetById(id uint64) (domain.MovieSelection, error) {
	for _, movieSelection := range repo.storage {
		if movieSelection.ID == id {
			return movieSelection, nil
		}
	}
	return domain.MovieSelection{}, ErrMovieSelectionNotFound
}
