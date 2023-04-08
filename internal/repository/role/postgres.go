package role

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

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (map[uint64]domain.Role, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	result := map[uint64]domain.Role{}
	for rows.Next() {
		var id uint64
		r := domain.Role{}
		err = rows.Scan(&id, &r.ID, &r.Title)
		result[id] = r
	}
	return result, nil
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) (map[uint64]domain.Role, error) {
	query := `select crp.person_id, r.id, r.title from roles r 
    		  join content_roles_persons crp on r.id = crp.role_id
    		  where crp.content_id = $1`
	return repo.fetch(ctx, query, ContentID)
}
