package genre

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
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
	query := strings.Join([]string{fetchQueryTemplate, filterByIDs}, " ")
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
	joinOnPerson := `join content_roles_persons crp on c.id = crp.content_id
					 join persons p on crp.person_id = p.id`
	filterByPersonID := `where p.id = $1`
	query := strings.Join([]string{fetchQueryTemplate, joinOnPerson, filterByPersonID}, " ")
	genres, err := repo.fetch(ctx, query, PersonID)
	if err != nil {
		return nil, err
	}
	uniqueGenres := map[domain.Genre]struct{}{}
	for _, genreSlice := range genres {
		for _, genre := range genreSlice {
			uniqueGenres[genre] = struct{}{}
		}
	}
	result := make([]domain.Genre, 0, len(uniqueGenres))
	for genre := range uniqueGenres {
		result = append(result, genre)
	}
	return result, nil
}
