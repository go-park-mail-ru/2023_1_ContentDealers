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

func (repo *Repository) Add(movies domain.Film) {
}

func (repo *Repository) GetByID(id uint64) (domain.Film, error) {
	movie := domain.Film{}
	err := repo.DB.
		QueryRow(`select f.id, title, description, preview_url from films f
                        join content c on f.content_id = c.id where f.id = $1`, id).
		Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PreviewURL)
	if err != nil {
		return domain.Film{}, err
	}
	return movie, nil
}

func (repo *Repository) GetAll() ([]domain.Film, error) {
	var movies []domain.Film
	rows, err := repo.DB.Query("SELECT id, title, description, preview_url FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		movie := domain.Film{}
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PreviewURL)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
