package country

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

const fetchQueryTemplate = `select countries.id, countries.name from countries 
    						join content_countries cc on countries.id = cc.country_id
							join content c on cc.content_id = c.id`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Country, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Country
	for rows.Next() {
		c := domain.Country{}
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Country, error) {
	filterByContentID := `where c.id = $1`
	orderByID := `order by countries.id`
	query := strings.Join([]string{fetchQueryTemplate, filterByContentID, orderByID}, " ")
	return repo.fetch(ctx, query, ContentID)
}
