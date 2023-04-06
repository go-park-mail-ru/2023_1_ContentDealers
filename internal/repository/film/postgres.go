package film

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

const (
	queryFetchTemplate = ``
)

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Film, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []domain.Film
	idToIndex := map[uint64]int{}
	for rows.Next() {
		film := domain.Film{}
		personRole := domain.ContentPerson{}
		genre := domain.Genre{}
		selection := domain.ContentSelection{}
		country := domain.Country{}
		err = rows.Scan(&film.ID, &film.ContentURL, &film.Content.ID, &film.Title, &film.Description, &film.Rating,
			&film.Year, &film.IsFree, &film.AgeLimit, &film.TrailerURL, &film.PreviewURL, &film.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, film)
	}
}

func (repo *Repository) Add(movies domain.Film) {
}

func (repo *Repository) GetByID(id uint64) (domain.Film, error) {
	film := domain.Film{}
	err := repo.DB.
		QueryRow(`select f.id, f.content_url, c.id, c.title, c.description, 
       					c.rating, c.year, c.is_free, c.age_limit, c.preview_url, 
       					c.trailer_url, c.type from films f 
       					join content c on c.id = f.content_id
       					left join content_roles_persons crp on c.id = crp.content_id
       					join roles r on crp.role_id = r.id
       					join persons p on crp.person_id = p.id
       					left join content_genres cg on c.id = cg.content_id
       					join genres g on cg.genre_id = g.id
       					left join content_selections cs on c.id = cs.content_id
       					join selections s on cs.selection_id = s.id`, id).
		Scan(&film.ID, &film.Title, &film.Description, &film.PreviewURL)
	if err != nil {
		return domain.Film{}, err
	}
	return film, nil
}

func (repo *Repository) GetAll() ([]domain.Film, error) {
	var movies []domain.Film
	rows, err := repo.DB.Query("SELECT id, title, description, preview_url FROM movies join ")
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
