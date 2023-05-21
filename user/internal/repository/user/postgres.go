package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	"github.com/google/uuid"
)

const (
	dirAvatars             = "media/avatars"
	allPerms   os.FileMode = 0777
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

func (repo *Repository) Add(ctx context.Context, user domain.User) (domain.User, error) {
	var lastInsertedID uint64

	err := repo.DB.QueryRowContext(ctx,
		`insert into users (email, password_hash, avatar_url) 
        values ($1, $2, $3) 
        returning id`,
		user.Email,
		user.PasswordHash,
		user.AvatarURL,
	).Scan(&lastInsertedID)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}
	user.ID = lastInsertedID
	return user, nil
}

func (repo *Repository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.
		QueryRowContext(ctx,
			`select id, email, password_hash, avatar_url, sub_expiration FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.AvatarURL, &user.SubscriptionExpiryDate)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.
		QueryRowContext(ctx,
			`select id, email, password_hash, avatar_url, sub_expiration FROM users WHERE id = $1`, id).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.AvatarURL, &user.SubscriptionExpiryDate)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func getDirByDate(date time.Time) string {
	year := fmt.Sprintf("%d", date.Year())
	month := fmt.Sprintf("%d", int(date.Month()))
	day := fmt.Sprintf("%d", date.Day())

	// TODO: разделитель пути зависит от операционной системы
	return fmt.Sprintf("%s/%s/%s", year, month, day)
}

func (repo *Repository) DeleteAvatar(ctx context.Context, user domain.User) error {
	if user.AvatarURL == "" || user.AvatarURL == "media/avatars/default_avatar.jpg" {
		repo.logger.WithRequestID(ctx).Trace("delete avatar, but it is empty")
		return nil
	}

	if _, err := os.Stat(user.AvatarURL); !errors.Is(err, os.ErrNotExist) {
		err = os.Remove(user.AvatarURL)
		if err != nil {
			repo.logger.WithRequestID(ctx).Trace(err)
			return err
		}
	}

	_, err := repo.DB.ExecContext(ctx,
		`update users 
		set avatar_url = 'media/avatars/default_avatar.jpg'
		where id = $1;`,
		user.ID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (repo *Repository) UpdateAvatar(ctx context.Context, user domain.User, file io.Reader) (domain.User, error) {
	err := repo.DeleteAvatar(ctx, user)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}

	dir := fmt.Sprintf("%s/%s", dirAvatars, getDirByDate(time.Now()))

	// TODO: хардкод прав на директорию
	err = os.MkdirAll(dir, allPerms)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, fmt.Errorf("failed to create folder for avatar")
	}

	avatarUUID := uuid.New().String()

	filepath := fmt.Sprintf("%s/%d_%s.jpg", dir, user.ID, avatarUUID)

	outAvatar, err := os.Create(filepath)
	defer outAvatar.Close()
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, fmt.Errorf("failed to create avatar file")
	}
	_, err = io.Copy(outAvatar, file)

	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, fmt.Errorf("failed to copy avatar file to local directory")
	}

	_, err = repo.DB.ExecContext(ctx,
		`update users 
		set avatar_url = $1
		where id = $2;`,
		filepath,
		user.ID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return domain.User{}, err
	}
	user.AvatarURL = filepath
	return user, nil
}

func (repo *Repository) Update(ctx context.Context, user domain.User) error {
	_, err := repo.DB.ExecContext(ctx,
		`update users 
		set email = $1,
			password_hash = $2
		where id = $3;`,
		user.Email,
		user.PasswordHash,
		user.ID,
	)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

func (repo *Repository) Subscribe(ctx context.Context, user domain.User) error {
	_, err := repo.DB.ExecContext(ctx,
		`update users set sub_expiration = greatest(current_date, sub_expiration) + interval '1 month'
				where id = $1`, user.ID)
	if err != nil {
		repo.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}
