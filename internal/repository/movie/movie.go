package movie

import (
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (repo *Repository) Add(movies domain.Movie) {
}

func (repo *Repository) GetByID(id uint64) (domain.Movie, error) {
	movie := domain.Movie{}
	err := repo.DB.
		QueryRow(`SELECT id, title, description, preview_url FROM movies WHERE id = $1`, id).
		Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PreviewURL)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (repo *Repository) GetAll() ([]domain.Movie, error) {
	movies := []domain.Movie{}
	rows, err := repo.DB.Query("SELECT id, title, description, preview_url FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		movie := domain.Movie{}
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PreviewURL)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
