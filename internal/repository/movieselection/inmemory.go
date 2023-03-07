package movieselection

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

var _ contract.MovieSelectionRepository = (*InMemoryRepository)(nil)

type InMemoryRepository struct {
	storage []domain.MovieSelection
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{}
}

func (repo *InMemoryRepository) Add(selections domain.MovieSelection) {
	repo.storage = append(repo.storage, selections)
}

func (repo *InMemoryRepository) GetAll() ([]domain.MovieSelection, error) {
	return repo.storage, nil
}

func (repo *InMemoryRepository) GetByID(id uint64) (domain.MovieSelection, error) {
	for _, movieSelection := range repo.storage {
		if movieSelection.ID == id {
			return movieSelection, nil
		}
	}
	return domain.MovieSelection{}, domain.ErrMovieSelectionNotFound
}
