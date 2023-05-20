package person

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/sharederrors"
)

const searchLimit = 6

type Repository struct {
	DB           *sql.DB
	simThreshold float32
}

func NewRepository(db *sql.DB, simThreshold float32) Repository {
	return Repository{DB: db, simThreshold: simThreshold}
}

const fetchQueryTemplate = `select p.id, p.name, p.gender, p.growth, p.birthplace, p.avatar_url, p.age from persons p`

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) ([]domain.Person, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Person{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var result []domain.Person
	for rows.Next() {
		p := domain.Person{}
		err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.Growth, &p.Birthplace, &p.AvatarURL, &p.Age)
		if err != nil {
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
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Person{}, sharederrors.ErrRepoNotFound
		}
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

func (repo *Repository) Search(ctx context.Context, query domain.SearchQuery) (domain.SearchPerson, error) {
	likeQuery := "%" + query.Query + "%"
	fullQuery := `select s.id, s.name, s.gender, s.growth, s.birthplace, s.avatar_url, s.age from (
				(select id, 1 sim, name, gender, growth, birthplace, avatar_url, age from persons
				 where lower(name) like $1)
				union all
				(select id, SIMILARITY($2, name) sim, name, gender, growth, birthplace, avatar_url, age 
					from persons
					where SIMILARITY($2, name) > $3)
				) s
				group by s.id, s.name, s.gender, s.growth, s.birthplace, s.avatar_url, s.age
				order by max(s.sim) desc
				limit $4 offset $5;`

	rows, err := repo.DB.QueryContext(ctx, fullQuery, likeQuery, repo.simThreshold, query.Limit, query.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.SearchPerson{}, nil
		}
		return domain.SearchPerson{}, err
	}
	defer rows.Close()

	var result domain.SearchPerson
	for rows.Next() {
		p := domain.Person{}
		err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.Growth, &p.Birthplace, &p.AvatarURL, &p.Age)
		if err != nil {
			return domain.SearchPerson{}, err
		}
		result.Persons = append(result.Persons, p)
	}

	row := repo.DB.QueryRowContext(ctx, `select count(*) from persons`)
	err = row.Scan(&result.Total)
	return result, err
}
