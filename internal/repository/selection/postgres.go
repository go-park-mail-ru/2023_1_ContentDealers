package selection

import (
	"context"
	"database/sql"
	"strings"

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

const queryFetchTemplate = `select s.id, s.title from selections s`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Selection, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Selection
	for rows.Next() {
		s := domain.Selection{}
		err = rows.Scan(&s.ID, &s.Title)
		if err != nil {
			repo.logger.WithFields(logrus.Fields{
				"request_id": ctx.Value("requestID").(string),
			}).Trace(err)
			return nil, err
		}
		result = append(result, s)
	}
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
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
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
