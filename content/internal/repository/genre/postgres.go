package genre

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/sharederrors"
	"github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

const fetchQueryTemplate = `select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
							join content c on cg.content_id = c.id`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (map[uint64][]domain.Genre, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return map[uint64][]domain.Genre{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	result := map[uint64][]domain.Genre{}
	for rows.Next() {
		var contentID uint64
		g := domain.Genre{}
		err = rows.Scan(&contentID, &g.ID, &g.Name)
		if err != nil {
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
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Genre{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Genre
	for rows.Next() {
		g := domain.Genre{}
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, g)
	}
	return result, nil
}

func (repo *Repository) GetAll(ctx context.Context) ([]domain.Genre, error) {
	query := `select id, name from genres order by id;`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharederrors.ErrRepoNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Genre
	for rows.Next() {
		g := domain.Genre{}
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, g)
	}
	return result, nil
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Genre, error) {
	query := `select id, name from genres where id = $1;`
	row := repo.DB.QueryRowContext(ctx, query, id)

	result := domain.Genre{}
	err := row.Scan(&result.ID, &result.Name)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return domain.Genre{}, sharederrors.ErrRepoNotFound
	}
	return result, err
}
