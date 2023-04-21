package role

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (map[uint64]domain.Role, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return nil, err
	}
	defer rows.Close()

	result := map[uint64]domain.Role{}
	for rows.Next() {
		var id uint64
		r := domain.Role{}
		err = rows.Scan(&id, &r.ID, &r.Title)
		if err != nil {
			repo.logger.WithFields(logrus.Fields{
				"request_id": ctx.Value("requestID").(string),
			}).Trace(err)
			return nil, err
		}
		result[id] = r
	}
	return result, nil
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) (map[uint64]domain.Role, error) {
	query := `select crp.person_id, r.id, r.title from roles r 
    		  join content_roles_persons crp on r.id = crp.role_id
    		  where crp.content_id = $1
			  order by r.id`
	return repo.fetch(ctx, query, ContentID)
}

func (repo *Repository) GetByPersonID(ctx context.Context, PersonID uint64) ([]domain.Role, error) {
	query := `select distinct(r.id), r.title from roles r 
    		  join content_roles_persons crp on r.id = crp.role_id
    		  where crp.person_id = $1
			  order by r.id`
	rows, err := repo.DB.QueryContext(ctx, query, PersonID)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Role
	for rows.Next() {
		r := domain.Role{}
		err = rows.Scan(&r.ID, &r.Title)
		if err != nil {
			repo.logger.WithFields(logrus.Fields{
				"request_id": ctx.Value("requestID").(string),
			}).Trace(err)
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
