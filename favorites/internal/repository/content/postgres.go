package content

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
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

func (repo *Repository) Get(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error) {
	var orderDate string
	if options.Order == "old" {
		orderDate = "ask"
	} else {
		orderDate = "desc"
	}
	rows, err := repo.DB.QueryContext(ctx,
		`select user_id, content_id, created_at 
		from users_content_favorites 
		where user_id = $1
		order by created_at &2;`,
		options.UserID,
		orderDate,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.FavoritesContent{}, nil
		}
		return domain.FavoritesContent{}, err
	}
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
	return result, nil
}
