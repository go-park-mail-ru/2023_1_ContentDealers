package selection

import (
	"context"
	"database/sql"
	"fmt"
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
	defer rows.Close()

	var result []domain.Selection
	for rows.Next() {
		s := domain.Selection{}
		err = rows.Scan(&s.ID, &s.Title)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	fmt.Println(query, result)
	return result, nil
}

func (repo *Repository) GetAll(ctx context.Context, limit uint, offset uint) ([]domain.Selection, error) {
	orderByID := `order by id desc`
	limitAndOffset := `limit $1 offset $2`
	query := strings.Join([]string{queryFetchTemplate, orderByID, limitAndOffset}, " ")
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
	orderByID := `order by s.id`
	query := strings.Join([]string{queryFetchTemplate, joinContent, filterByContentID, orderByID}, " ")
	return repo.fetch(ctx, query, ContentID)
}
