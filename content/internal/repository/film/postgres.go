package film

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
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
		repo.logger.Trace(err)
	}
	return film, err
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	filterByContentID := `where content_id = $1`
	query := strings.Join([]string{fetchQueryTemplate, filterByContentID}, " ")
	return repo.fetch(ctx, query, ContentID)
}
