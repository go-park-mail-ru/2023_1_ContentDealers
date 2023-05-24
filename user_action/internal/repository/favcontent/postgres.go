package favcontent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

func (repo *Repository) Delete(ctx context.Context, favorite domain.FavoriteContent) error {
	_, err := repo.DB.ExecContext(ctx,
		`delete from users_content_favorites 
		where user_id = $1 and
		content_id = $2;`,
		favorite.UserID,
		favorite.ContentID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
	}
	return err
}

func (repo *Repository) Add(ctx context.Context, favorite domain.FavoriteContent) error {
	_, err := repo.DB.ExecContext(ctx,
		`insert into users_content_favorites 
		(user_id, content_id) 
		values ($1, $2);`,
		favorite.UserID,
		favorite.ContentID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
	}
	return err
}

func (repo *Repository) HasFav(ctx context.Context, favorite domain.FavoriteContent) (bool, error) {
	var count int
	err := repo.DB.QueryRowContext(ctx,
		`select count(*)
		from users_content_favorites 
		where user_id = $1 and content_id = $2`,
		favorite.UserID, favorite.ContentID).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 1 {
		return true, nil
	}
	return false, nil
}

func (repo *Repository) Get(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error) {
	var sortDate string
	if options.SortDate == "old" {
		sortDate = "asc"
	} else {
		sortDate = "desc"
	}

	limit := options.Limit
	offset := options.Offset

	query := `select user_id, content_id, created_at 
			from users_content_favorites 
			where user_id = $1`
	limitOffset := `limit $2 offset $3`

	rows, err := repo.DB.QueryContext(ctx,
		fmt.Sprintf("%s order by created_at %s %s;", query, sortDate, limitOffset),
		options.UserID, limit+1, offset,
	)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.FavoritesContent{}, nil
		}
		return domain.FavoritesContent{}, err
	}

	defer rows.Close()

	result := domain.FavoritesContent{}
	for rows.Next() {
		fav := domain.FavoriteContent{}
		err := rows.Scan(&fav.UserID, &fav.ContentID, &fav.DateAdding)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return domain.FavoritesContent{}, err
		}
		result.Favorites = append(result.Favorites, fav)
	}

	result.IsLast = true
	if len(result.Favorites) == int(limit+1) {
		result.Favorites = result.Favorites[:len(result.Favorites)-1]
		result.IsLast = false
	}

	return result, nil
}
