package rating

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func (repo *Repository) Delete(ctx context.Context, rating domain.Rating) error {
	_, err := repo.DB.ExecContext(ctx,
		`delete from ratings 
		where user_id = $1 and
		content_id = $2;`,
		rating.UserID,
		rating.ContentID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
	}
	return err
}

// numeric(4,2) <-> float32 ???

func (repo *Repository) Add(ctx context.Context, rating domain.Rating) error {
	_, err := repo.DB.ExecContext(ctx,
		`insert into ratings 
		(user_id, content_id, rating) 
		values ($1, $2, $3);`,
		rating.UserID,
		rating.ContentID,
		rating.Rating,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
	}
	return err
}

func (repo *Repository) Has(ctx context.Context, rating domain.Rating) (domain.HasRating, error) {
	var ratingNum float32
	var createdAt time.Time
	err := repo.DB.QueryRowContext(ctx,
		`select rating, created_at
		from ratings 
		where user_id = $1 and content_id = $2`,
		rating.UserID, rating.ContentID).Scan(&ratingNum, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.HasRating{HasRating: false}, nil
		}
		return domain.HasRating{}, err
	}
	rating.Rating = ratingNum
	rating.DateAdding = createdAt
	hasRating := domain.HasRating{
		HasRating: true,
		Rating:    rating,
	}
	return hasRating, nil
}

func (repo *Repository) GetByUser(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	var sortDate string
	if options.SortDate == "old" {
		sortDate = "asc"
	} else {
		sortDate = "desc"
	}

	limit := options.Limit
	offset := options.Offset

	query := `select user_id, content_id, rating, created_at
			from ratings 
			where user_id = $1`
	limitOffset := `limit $2 offset $3`

	rows, err := repo.DB.QueryContext(ctx,
		fmt.Sprintf("%s order by created_at %s %s;", query, sortDate, limitOffset),
		options.UserID, limit+1, offset,
	)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Ratings{}, nil
		}
		return domain.Ratings{}, err
	}
	result := domain.Ratings{}
	for rows.Next() {
		rate := domain.Rating{}
		err := rows.Scan(&rate.UserID, &rate.ContentID, &rate.Rating, &rate.DateAdding)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return domain.Ratings{}, err
		}
		result.Ratings = append(result.Ratings, rate)
	}

	result.IsLast = true
	if len(result.Ratings) == int(limit+1) {
		result.Ratings = result.Ratings[:len(result.Ratings)-1]
		result.IsLast = false
	}

	return result, nil
}

// нет сортировки, limit, offset
func (repo *Repository) GetByContent(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	rows, err := repo.DB.QueryContext(ctx,
		`select user_id, content_id, rating, created_at
		from ratings
		where content_id = $1`, options.ContentID)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Ratings{}, nil
		}
		return domain.Ratings{}, err
	}
	result := domain.Ratings{}
	for rows.Next() {
		rate := domain.Rating{}
		err := rows.Scan(&rate.UserID, &rate.ContentID, &rate.Rating, &rate.DateAdding)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return domain.Ratings{}, err
		}
		result.Ratings = append(result.Ratings, rate)
	}
	result.IsLast = true
	return result, nil
}
