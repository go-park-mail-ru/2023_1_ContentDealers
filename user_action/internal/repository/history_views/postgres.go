package history_views

import (
	"context"
	"database/sql"
	"database/sql/driver"
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

type Duration time.Duration

func (d Duration) Value() (driver.Value, error) {
	return driver.Value(int64(d)), nil
}

func (d *Duration) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case int64:
		*d = Duration(v)
	case nil:
		*d = Duration(0)
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.Duration from: %#v", v)
	}
	return nil
}

func (repo *Repository) UpdateProgressView(ctx context.Context, view domain.View) error {
	_, err := repo.DB.ExecContext(ctx,
		`insert into history_views (user_id, content_id, stop_view, duration)
			values ($1, $2, $3, $4)
			on conflict (user_id, content_id) do update
				set stop_view = excluded.stop_view, duration = excluded.duration;`,
		view.UserID,
		view.ContentID,
		Duration(view.StopView),
		Duration(view.Duration),
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
	}
	return err
}

func (repo *Repository) HasView(ctx context.Context, view domain.View) (domain.HasView, error) {
	var stopView Duration
	var duration Duration
	var createdAt time.Time
	err := repo.DB.QueryRowContext(ctx,
		`select stop_view, duration, created_at
		from history_views 
		where user_id = $1 and content_id = $2`,
		view.UserID, view.ContentID).Scan(&stopView, &duration, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.HasView{HasView: false}, nil
		}
		return domain.HasView{}, err
	}
	// view.StopView, _ = time.ParseDuration(fmt.Sprintf("%dns", stopView))
	view.StopView = time.Duration(stopView)
	view.Duration = time.Duration(duration)
	view.DateAdding = createdAt
	hasView := domain.HasView{
		HasView: true,
		View:    view,
	}
	return hasView, nil
}

func (repo *Repository) GetAllViewsByUser(ctx context.Context, options domain.ViewsOptions) (domain.Views, error) {
	var sortDate string
	if options.SortDate == "old" {
		sortDate = "asc"
	} else {
		sortDate = "desc"
	}

	limit := options.Limit
	offset := options.Offset

	query := `select user_id, content_id, stop_view, duration, created_at 
			from history_views 
			where user_id = $1`
	limitOffset := `limit $2 offset $3`

	rows, err := repo.DB.QueryContext(ctx,
		fmt.Sprintf("%s order by created_at %s %s;", query, sortDate, limitOffset),
		options.UserID, limit+1, offset,
	)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Views{}, nil
		}
		return domain.Views{}, err
	}
	result := domain.Views{}
	for rows.Next() {
		view := domain.View{}
		var stopView Duration
		var duration Duration
		err := rows.Scan(&view.UserID, &view.ContentID, &stopView, &duration, &view.DateAdding)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return domain.Views{}, err
		}

		view.StopView = time.Duration(stopView)
		view.Duration = time.Duration(duration)
		result.Views = append(result.Views, view)
	}

	result.IsLast = true
	if len(result.Views) == int(limit+1) {
		result.Views = result.Views[:len(result.Views)-1]
		result.IsLast = false
	}

	return result, nil
}

// FIXME: копипаст с доп условием в запросе...
func (repo *Repository) GetPartiallyViewsByUser(ctx context.Context, options domain.ViewsOptions) (domain.Views, error) {
	var sortDate string
	if options.SortDate == "old" {
		sortDate = "asc"
	} else {
		sortDate = "desc"
	}

	limit := options.Limit
	offset := options.Offset

	query := `select user_id, content_id, stop_view, duration, created_at 
			from history_views 
			where user_id = $1
			and stop_view::float / duration < 0.9`
	limitOffset := `limit $2 offset $3`

	rows, err := repo.DB.QueryContext(ctx,
		fmt.Sprintf("%s order by created_at %s %s;", query, sortDate, limitOffset),
		options.UserID, limit+1, offset,
	)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Views{}, nil
		}
		return domain.Views{}, err
	}
	result := domain.Views{}
	for rows.Next() {
		view := domain.View{}
		var stopView Duration
		var duration Duration
		err := rows.Scan(&view.UserID, &view.ContentID, &stopView, &duration, &view.DateAdding)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return domain.Views{}, err
		}

		view.StopView = time.Duration(stopView)
		view.Duration = time.Duration(duration)
		result.Views = append(result.Views, view)
	}

	result.IsLast = true
	if len(result.Views) == int(limit+1) {
		result.Views = result.Views[:len(result.Views)-1]
		result.IsLast = false
	}

	return result, nil
}
