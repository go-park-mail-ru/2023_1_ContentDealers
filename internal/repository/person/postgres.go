package person

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

const fetchQueryTemplate = `select id, name, gender, growth, birthplace, avatar_url, age from persons`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (domain.Person, error) {
	row := repo.DB.QueryRowContext(ctx, query, args...)
	p := domain.Person{}
	err := row.Scan(&p.ID, &p.Name, &p.Gender, &p.Growth, &p.Birthplace, &p.AvatarURL, &p.Age)
	return p, err
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	filterByIdQueryPart := `where p.id = $1`
	fullQuery := strings.Join([]string{fetchQueryTemplate, filterByIdQueryPart}, " ")
	return repo.fetch(ctx, fullQuery, id)
}
