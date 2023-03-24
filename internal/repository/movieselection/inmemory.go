package movieselection

import (
	"sync"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type InMemoryRepository struct {
	mu      sync.RWMutex
	storage []domain.MovieSelection
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{}
}

func (repo *InMemoryRepository) Add(selections domain.MovieSelection) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.storage = append(repo.storage, selections)
}

func (repo *InMemoryRepository) GetAll() ([]domain.MovieSelection, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
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
