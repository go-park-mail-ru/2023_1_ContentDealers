package person

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const searchLimit = 6

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

const fetchQueryTemplate = `select p.id, p.name, p.gender, p.growth, p.birthplace, p.avatar_url, p.age from persons p`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Person, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		repo.logger.Trace(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.Person
	for rows.Next() {
		p := domain.Person{}
		err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.Growth, &p.Birthplace, &p.AvatarURL, &p.Age)
		if err != nil {
			repo.logger.Trace(err)
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	filterByIDQueryPart := `where p.id = $1`
	orderByID := `order by p.id`
	fullQuery := strings.Join([]string{fetchQueryTemplate, filterByIDQueryPart, orderByID}, " ")
	persons, err := repo.fetch(ctx, fullQuery, id)
	if err != nil {
		repo.logger.Trace(err)
		return domain.Person{}, err
	}
	return persons[0], nil
}

func (repo *Repository) GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Person, error) {
	joinOnContent := `join content_roles_persons crp on crp.person_id = p.id`
	filterByContentID := `where crp.content_id = $1`
	orderByID := `order by p.id`
	query := strings.Join([]string{fetchQueryTemplate, joinOnContent, filterByContentID, orderByID}, " ")
	return repo.fetch(ctx, query, ContentID)
}

func (repo *Repository) Search(ctx context.Context, query string) ([]domain.Person, error) {
	likeQuery := "%" + query + "%"
	fullQuery := `select s.id, s.name, s.gender, s.growth, s.birthplace, s.avatar_url, s.age from (
				(select id, 1 sim, name, gender, growth, birthplace, avatar_url, age from persons
				 where lower(name) like $1)
				union all
				(select id, SIMILARITY($2, name) sim, name, gender, growth, birthplace, avatar_url, age from persons)
				) s
				group by s.id, s.name, s.gender, s.growth, s.birthplace, s.avatar_url, s.age
				order by max(s.sim) desc
				limit $3;`

	return repo.fetch(ctx, fullQuery, likeQuery, query, searchLimit)
}
