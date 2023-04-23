package content

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/lib/pq"
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

const fetchQueryTemplate = `select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       						c.trailer_url, c.preview_url, c.type from content c`

func (repo *Repository) fetchByIDs(ctx context.Context, query string, IDs []uint64) ([]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx, query, pq.Array(IDs))
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.Year, &c.IsFree, &c.AgeLimit, &c.TrailerURL,
			&c.PreviewURL, &c.Type)
		if err != nil {
			repo.logger.Trace(err)
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
		repo.logger.Trace(err)
		return domain.Content{}, err
	}
	return content[0], nil
}

func (repo *Repository) GetBySelectionIDs(ctx context.Context, IDs []uint64) (map[uint64][]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx,
		`select cs.selection_id, c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       		   c.trailer_url, c.preview_url, c.type from content c 
       		   join content_selections cs on c.id = cs.content_id
       		   where cs.selection_id = any($1)
			   order by c.rating desc`, pq.Array(IDs))
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	result := map[uint64][]domain.Content{}
	for rows.Next() {
		var selectionID uint64
		c := domain.Content{}
		err = rows.Scan(&selectionID, &c.ID, &c.Title, &c.Description, &c.Rating, &c.Year, &c.IsFree, &c.AgeLimit,
			&c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			repo.logger.Trace(err)
			return nil, err
		}
		result[selectionID] = append(result[selectionID], c)
	}
	return result, err
}

func (repo *Repository) GetByPersonID(ctx context.Context, id uint64) ([]domain.Content, error) {
	rows, err := repo.DB.QueryContext(ctx,
		`select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       		   c.trailer_url, c.preview_url, c.type from content c 
       		   join content_roles_persons crp on c.id = crp.content_id
       		   where crp.person_id = $1
			   order by c.rating desc`, id)
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Content
	for rows.Next() {
		c := domain.Content{}
		err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.Year, &c.IsFree, &c.AgeLimit,
			&c.TrailerURL, &c.PreviewURL, &c.Type)
		if err != nil {
			repo.logger.Trace(err)
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}
