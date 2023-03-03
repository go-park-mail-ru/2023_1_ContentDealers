package repository

import (
	"errors"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

var ErrMovieNotFound = errors.New("movie not found")

var _ usecase.MovieRepository = (*MovieInMemoryRepository)(nil)

type MovieInMemoryRepository struct {
	storage []domain.Movie
}

func NewMovieInMemoryRepository() MovieInMemoryRepository {
	return MovieInMemoryRepository{}
}

func (repo *MovieInMemoryRepository) Add(movies domain.Movie) {
	repo.storage = append(repo.storage, movies)
}

func (repo *MovieInMemoryRepository) GetById(id uint64) (domain.Movie, error) {
	for _, movie := range repo.storage {
		if movie.ID == id {
			return movie, nil
		}
	}
	return domain.Movie{}, ErrMovieNotFound
}

func (repo *MovieInMemoryRepository) GetAll() ([]domain.Movie, error) {
	return repo.storage, nil
}
