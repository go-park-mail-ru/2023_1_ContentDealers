package movie

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

var _ contract.MovieRepository = (*InMemoryRepository)(nil)

type InMemoryRepository struct {
	storage []domain.Movie
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{}
}

func (repo *InMemoryRepository) Add(movies domain.Movie) {
	repo.storage = append(repo.storage, movies)
}

func (repo *InMemoryRepository) GetByID(id uint64) (domain.Movie, error) {
	for _, movie := range repo.storage {
		if movie.ID == id {
			return movie, nil
		}
	}
	return domain.Movie{}, domain.ErrMovieNotFound
}

func (repo *InMemoryRepository) GetAll() ([]domain.Movie, error) {
	return repo.storage, nil
}
