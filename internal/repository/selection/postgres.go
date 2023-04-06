package selection

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

const queryFetchTemplate = `select s.id, s.title from selections s`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Selection, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var result []domain.Selection
	for rows.Next() {
		s := domain.Selection{}
		err = rows.Scan(&s.ID, &s.Title)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (repo *Repository) GetAll(ctx context.Context, limit uint, offset uint) ([]domain.Selection, error) {
	orderById := `order by id desc`
	limitAndOffset := `limit $1 offset $2`
	query := strings.Join([]string{queryFetchTemplate, orderById, limitAndOffset}, " ")
	return repo.fetch(ctx, query, limit, offset)
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Selection, error) {
	filterByID := `where id = $1`
	query := strings.Join([]string{queryFetchTemplate, filterByID}, " ")
	selections, err := repo.fetch(ctx, query, id)
	if err != nil {
		return domain.Selection{}, err
	}
	return selections[0], nil
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Selection, error) {
	joinContent := `join content_selections cs on cs.selection_id = s.id
					join content c on c.id = cs.content_id`
	filterByContentID := `where c.id = $1`
	query := strings.Join([]string{queryFetchTemplate, joinContent, filterByContentID}, " ")
	return repo.fetch(ctx, query, ContentID)
}
