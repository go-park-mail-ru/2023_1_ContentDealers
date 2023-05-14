package content

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/sharederrors"
	"github.com/lib/pq"
)

const searchLimit = 6
const simThreshold = 0.14

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

const fetchQueryTemplate = `select c.id, c.title, c.description, c.rating, c.sum_ratings, c.count_ratings,
 							c.year, c.is_free, c.age_limit,
       						c.trailer_url, c.preview_url, c.type from content c`

func (repo *Repository) fetchByIDs(ctx context.Context, query string, IDs []uint64) ([]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx, query, pq.Array(IDs))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharederrors.ErrRepoNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.SumRatings, &c.CountRatings,
			&c.Year, &c.IsFree, &c.AgeLimit, &c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (repo *Repository) GetByIDs(ctx context.Context, ids []uint64) ([]domain.Content, error) {
	filterByIDs := `where c.id = any($1)`
	query := strings.Join([]string{fetchQueryTemplate, filterByIDs}, " ")
	return repo.fetchByIDs(ctx, query, ids)
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Content, error) {
	content, err := repo.GetByIDs(ctx, []uint64{id})
	if err != nil {
		return domain.Content{}, err
	}
	return content[0], nil
}

func (repo *Repository) GetBySelectionIDs(ctx context.Context, IDs []uint64) (map[uint64][]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx,
		`select cs.selection_id, c.id, c.title, c.description, c.rating, c.sum_ratings, c.count_ratings,
       		   c.year, c.is_free, c.age_limit,
       		   c.trailer_url, c.preview_url, c.type from content c 
       		   join content_selections cs on c.id = cs.content_id
       		   where cs.selection_id = any($1)
			   order by c.rating desc`, pq.Array(IDs))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return map[uint64][]domain.Content{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	result := map[uint64][]domain.Content{}
	for rows.Next() {
		var selectionID uint64
		c := domain.Content{}
		err = rows.Scan(&selectionID, &c.ID, &c.Title, &c.Description, &c.Rating, &c.SumRatings, &c.CountRatings,
			&c.Year, &c.IsFree, &c.AgeLimit, &c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			return nil, err
		}
		result[selectionID] = append(result[selectionID], c)
	}
	return result, err
}

func (repo *Repository) GetByPersonID(ctx context.Context, id uint64) ([]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx,
		`select c.id, c.title, c.description, c.rating, c.sum_ratings, c.count_ratings,
       		   c.year, c.is_free, c.age_limit,
       		   c.trailer_url, c.preview_url, c.type from content c 
       		   join content_roles_persons crp on c.id = crp.content_id
       		   where crp.person_id = $1
			   order by c.rating desc`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Content{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.SumRatings, &c.CountRatings,
			&c.Year, &c.IsFree, &c.AgeLimit, &c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (repo *Repository) GetByGenreOptions(ctx context.Context, options domain.ContentFilter) ([]domain.Content, error) {
	joinGenres := `join content_genres cg on cg.content_id = c.id
                   join genres g on cg.genre_id = g.id`
	filterByGenreID := `where g.id = $1`
	orderByRating := `order by c.rating desc`
	limitOffset := `limit $2 offset $3;`
	query := strings.Join([]string{fetchQueryTemplate, joinGenres, filterByGenreID, orderByRating, limitOffset}, " ")
	rows, err := repo.DB.QueryContext(ctx, query, options.ID, options.Limit, options.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Content{}, nil
		}
		return nil, err
	}
	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.SumRatings, &c.CountRatings,
			&c.Year, &c.IsFree, &c.AgeLimit, &c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (repo *Repository) Search(ctx context.Context, query string) ([]domain.Content, error) {
	likeQuery := "%" + query + "%"
	rows, err := repo.DB.QueryContext(ctx,
		`select s.id, s.title, s.description, s.rating, s.sum_ratings, s.count_ratings,
       			s.year, s.is_free, s.age_limit,
       			s.trailer_url, s.preview_url, s.type from (
				(select id, 1 sim, title, description, rating, sum_ratings, count_ratings, year, is_free, age_limit,
					trailer_url, preview_url, type 
					from content
				 	where lower(title) like $1)
					
				union all
				
				(select id, SIMILARITY($2, title) sim, title, description, rating, sum_ratings, count_ratings,
					year, is_free, age_limit, trailer_url, preview_url, type 
					from content
					where SIMILARITY($2, title) > $3)
				) s

				group by s.id, s.title, s.description, s.rating, s.sum_ratings, s.count_ratings, s.year, s.is_free, s.age_limit,
				s.trailer_url, s.preview_url, s.type
				order by max(s.sim) desc, s.rating desc
				limit $4;`, likeQuery, query, simThreshold, searchLimit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Content{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.SumRatings, &c.CountRatings, &c.Year, &c.IsFree,
			&c.AgeLimit, &c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (repo *Repository) GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	query := `select id, content_url from films where content_id = $1`
	row := repo.DB.QueryRowContext(ctx, query, ContentID)
	film := domain.Film{}
	err := row.Scan(&film.ID, &film.ContentURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Film{}, sharederrors.ErrRepoNotFound
		}
		return domain.Film{}, err
	}
	return film, err
}

func (repo *Repository) GetSeriesByContentID(ctx context.Context, ContentID uint64) (domain.Series, error) {
	query := `select s.id, e.id, e.season_num, e.episode_num, e.content_url, e.preview_url, e.release_date, e.title
		 from series s join episodes e on s.id = e.series_id
         where s.content_id = $1 order by season_num, episode_num`
	rows, err := repo.DB.QueryContext(ctx, query, ContentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Series{}, sharederrors.ErrRepoNotFound
		}
		return domain.Series{}, err
	}
	defer rows.Close()

	var s domain.Series
	for rows.Next() {
		e := domain.Episode{}
		err = rows.Scan(&s.ID, &e.ID, &e.SeasonNum, &e.EpisodeNum, &e.ContentURL, &e.PreviewURL,
			&e.ReleaseDate, &e.Title)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return domain.Series{}, sharederrors.ErrRepoNotFound
			}
			return domain.Series{}, err
		}
		s.Episodes = append(s.Episodes, e)
	}
	return s, err
}

func (repo *Repository) AddRating(ctx context.Context, ContentID uint64, rating float32) error {
	_, err := repo.DB.ExecContext(ctx,
		`UPDATE content 
		SET sum_ratings = sum_ratings + $1, count_ratings = count_ratings + 1 
		WHERE id = $2;`,
		rating, ContentID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteRating(ctx context.Context, ContentID uint64, rating float32) error {
	_, err := repo.DB.ExecContext(ctx,
		`UPDATE content 
		SET sum_ratings = sum_ratings - $1, count_ratings = count_ratings - 1 
		WHERE id = $2;`,
		rating, ContentID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
