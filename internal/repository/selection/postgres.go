package selection

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

func (repo *Repository) Add(selections domain.Selection) {
}

func (repo *Repository) GetAll() ([]domain.Selection, error) {
	var selections []domain.Selection
	rows, err := repo.DB.Query(
		`select s.id, s.title, m.id, m.title, m.description, m.preview_url 
		FROM selections s
		join movie_selections ms on ms.selection_id = s.id
		join movies m on ms.movie_id = m.id
		order by s.id, m.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		s := domain.Selection{}
		m := domain.Film{}
		err = rows.Scan(&s.ID, &s.Title, &m.ID, &m.Title, &m.Description, &m.PreviewURL)
		if err != nil {
			return nil, err
		}
		if len(selections) == 0 || selections[len(selections)-1].ID != s.ID {
			s.Movies = append(s.Movies, &m)
			selections = append(selections, s)
		} else {
			lastSelection := &selections[len(selections)-1].Movies
			*lastSelection = append(*lastSelection, &m)
		}
	}
	return selections, nil
}

func (repo *Repository) GetByID(id uint64) (domain.Selection, error) {
	rows, err := repo.DB.Query(
		`select s.id, s.title, m.id, m.title, m.description, m.preview_url 
			FROM selections s
			join movie_selections ms on ms.selection_id = s.id
			join movies m on ms.movie_id = m.id
			where s.id = $1
			order by s.id, m.id`, id)
	if err != nil {
		return domain.Selection{}, err
	}
	defer rows.Close()
	selection := domain.Selection{}
	i := 0
	for rows.Next() {
		// TODO: возможнос стоит сделать два запроса, один для
		// названия выборки, второй для всех входящих в нее фильмов
		s := domain.Selection{}
		m := domain.Film{}
		if i == 0 {
			err = rows.Scan(&selection.ID, &selection.Title, &m.ID,
				&m.Title, &m.Description, &m.PreviewURL)
		} else {
			err = rows.Scan(&s.ID, &s.Title, &m.ID, &m.Title,
				&m.Description, &m.PreviewURL)
		}
		if err != nil {
			return domain.Selection{}, err
		}
		selection.Movies = append(selection.Movies, &m)
		i += 1
	}
	return selection, nil
}
