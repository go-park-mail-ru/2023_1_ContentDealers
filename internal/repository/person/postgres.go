package person

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const (
	fetchOneQueryTemplate = `select p.id, p.name, p.gender, p.growth, p.birthplace, p.avatar_url, p.age,
						c.id, c.title, r.id, r.title, g.id, g.name from persons p
    					left join content_roles_persons crp on p.id = crp.person_id
    					join roles r on crp.role_id = r.id
    					join content c on crp.content_id = c.id
    					left join content_genres cg on c.id = cg.content_id
    					join genres g on cg.genre_id = g.id`
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (repo *Repository) fetch(ctx context.Context, query string, args ...any) (domain.Person, error) {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return domain.Person{}, err
	}
	defer rows.Close()

	p := domain.Person{}
	participatedInAll := map[domain.PersonContent]struct{}{}
	roleAll := map[domain.Role]struct{}{}
	genreAll := map[domain.Genre]struct{}{}
	for rows.Next() {
		participatedIn := domain.PersonContent{}
		role := domain.Role{}
		genre := domain.Genre{}
		err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.Growth, &p.Birthplace, &p.AvatarURL, &p.Age,
			&participatedIn.ContentID, &participatedIn.Title, &role.ID, &role.Title, &genre.ID, &genre.Name)
		if err != nil {
			return p, err
		}
		participatedInAll[participatedIn] = struct{}{}
		roleAll[role] = struct{}{}
		genreAll[genre] = struct{}{}
	}
	for content := range participatedInAll {
		p.ParticipatedIn = append(p.ParticipatedIn, content)
	}
	for role := range roleAll {
		p.Roles = append(p.Roles, role)
	}
	for genre := range genreAll {
		p.Genres = append(p.Genres, genre)
	}
	return p, nil
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	filterByIdQueryPart := `where p.id = $1`
	fullQuery := strings.Join([]string{fetchOneQueryTemplate, filterByIdQueryPart}, " ")
	return repo.fetch(ctx, fullQuery, id)
}
