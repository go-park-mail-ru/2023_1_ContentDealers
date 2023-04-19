package genre

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
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

const fetchQueryTemplate = `select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
							join content c on cg.content_id = c.id`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (map[uint64][]domain.Genre, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	result := map[uint64][]domain.Genre{}
	for rows.Next() {
		var contentID uint64
		g := domain.Genre{}
		err = rows.Scan(&contentID, &g.ID, &g.Name)
		if err != nil {
			repo.logger.Trace(err)
			return nil, err
		}
		result[contentID] = append(result[contentID], g)
	}
	return result, nil
}

func (repo *Repository) GetByContentIDs(ctx context.Context, contentIDs []uint64) (map[uint64][]domain.Genre, error) {
	filterByIDs := `where c.id = any($1)`
	orderByID := `order by g.id`
	query := strings.Join([]string{fetchQueryTemplate, filterByIDs, orderByID}, " ")
	return repo.fetch(ctx, query, pq.Array(contentIDs))
}

func (repo *Repository) GetByContentID(ctx context.Context, contentID uint64) ([]domain.Genre, error) {
	ContentIDGenres, err := repo.GetByContentIDs(ctx, []uint64{contentID})
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	return ContentIDGenres[contentID], nil
}

func (repo *Repository) GetByPersonID(ctx context.Context, PersonID uint64) ([]domain.Genre, error) {
	query := `select distinct(g.id), g.name from genres g
			join content_genres cg on g.id = cg.genre_id
			join content_roles_persons crp on cg.content_id = crp.content_id
			where crp.person_id = $1
			order by g.id`
	rows, err := repo.DB.QueryContext(ctx, query, PersonID)
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Genre
	for rows.Next() {
		g := domain.Genre{}
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			repo.logger.Trace(err)
			return nil, err
		}
		result = append(result, g)
	}
	return result, nil
}
