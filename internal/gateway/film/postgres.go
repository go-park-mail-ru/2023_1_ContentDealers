package film

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

const fetchQueryTemplate = `select id, content_url from films`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (domain.Film, error) {
	row := repo.DB.QueryRowContext(ctx, query, args...)
	film := domain.Film{}
	err := row.Scan(&film.ID, &film.ContentURL)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
	}
	return film, err
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	filterByContentID := `where content_id = $1`
	query := strings.Join([]string{fetchQueryTemplate, filterByContentID}, " ")
	return repo.fetch(ctx, query, ContentID)
}