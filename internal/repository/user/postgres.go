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

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const (
	shortFormDate             = "2006-01-02"
	dirAvatars                = "media/avatars"
	allPerms      os.FileMode = 0777
)

type Repository struct {
	DB     *sql.DB
	logger logging.Logger
}

func NewRepository(db *sql.DB, logger logging.Logger) Repository {
	return Repository{DB: db, logger: logger}
}

func (repo *Repository) Add(ctx context.Context, user domain.User) (domain.User, error) {
	repo.logger.Tracef("input params Add(): %#v", user)
	var lastInsertedID uint64

	err := repo.DB.QueryRowContext(ctx,
		`insert into users (email, password_hash, date_birth, avatar_url) 
        values ($1, $2, $3, $4) 
        returning id`,
		user.Email,
		user.PasswordHash,
		user.DateBirth,
		user.AvatarURL,
	).Scan(&lastInsertedID)
	if err != nil {
		repo.logger.Trace(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}
	user.ID = lastInsertedID
	return user, nil
}

func (repo *Repository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	repo.logger.Tracef("input params GetByEmail(): %#v", email)
	user := domain.User{}
	err := repo.DB.
		QueryRowContext(ctx,
			`select id, email, password_hash, date_birth, avatar_url FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.DateBirth, &user.AvatarURL)
	if err != nil {
		repo.logger.Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (repo *Repository) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	// TODO: копипаст метода GetByEmail (нужен общий метод для запроса)
	repo.logger.Tracef("input params GetByID(): %#v", id)
	user := domain.User{}
	err := repo.DB.
		QueryRowContext(ctx,
			`select id, email, password_hash, date_birth, avatar_url FROM users WHERE id = $1`, id).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.DateBirth, &user.AvatarURL)
	if err != nil {
		repo.logger.Trace(err)
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
	repo.logger.Tracef("input params DeleteAvatar(): %#v", user)
	var updateDate time.Time
	// 1. удаление из локальной директории
	if user.AvatarURL == "" {
		repo.logger.Trace("delete avatar, but it is empty")
		return nil
	}
	err := repo.DB.
		QueryRowContext(ctx, `select updated_at FROM users WHERE id = $1`, user.ID).
		Scan(&updateDate)
	if err != nil {
		repo.logger.Trace(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrUserNotFound
		}
		return err
	}

	filepathForDelete := fmt.Sprintf("%s/%s/%d.jpg", dirAvatars, getDirByDate(updateDate), user.ID)
	// TODO: сомнительная логика
	if _, err = os.Stat(filepathForDelete); !errors.Is(err, os.ErrNotExist) {
		repo.logger.Trace(err)
		err = os.Remove(filepathForDelete)
		if err != nil {
			repo.logger.Trace(err)
			return err
		}
	}

	// 2. удаление урла из БД
	_, err = repo.DB.ExecContext(ctx,
		`update users 
		set avatar_url = 'media/avatars/default_avatar.jpg'
		where id = $1;`,
		user.ID,
	)
	if err != nil {
		repo.logger.Trace(err)
		return err
	}
	return nil
}

func (repo *Repository) UpdateAvatar(ctx context.Context, user domain.User, file io.Reader) (domain.User, error) {
	repo.logger.Tracef("input params UpdateAvatar(): %#v", user)
	err := repo.DeleteAvatar(ctx, user)

	if err != nil {
		repo.logger.Trace(err)
		return domain.User{}, err
	}

	// TODO: разделитель пути зависит от операционной системы
	dir := fmt.Sprintf("%s/%s", dirAvatars, getDirByDate(time.Now()))

	// если директория существовала, err == nil
	// TODO: хардкод прав на директорию
	err = os.MkdirAll(dir, allPerms)
	if err != nil {
		repo.logger.Trace(err)
		return domain.User{}, fmt.Errorf("failed to create folder for avatar")
	}
	filepath := fmt.Sprintf("%s/%d.jpg", dir, user.ID)

	outAvatar, err := os.Create(filepath)
	defer outAvatar.Close()
	if err != nil {
		repo.logger.Trace(err)
		return domain.User{}, fmt.Errorf("failed to create avatar file")
	}
	_, err = io.Copy(outAvatar, file)

	if err != nil {
		repo.logger.Trace(err)
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
		repo.logger.Trace(err)
		return domain.User{}, err
	}
	user.AvatarURL = filepath
	return user, nil
}

func (repo *Repository) Update(ctx context.Context, user domain.User) error {
	repo.logger.Tracef("input params Update(): %#v", user)
	// TODO: может поменять почту на уже существующую у др пользователя в системе, тогда возвращаем ошибку
	_, err := repo.DB.ExecContext(ctx,
		`update users 
		set email = $1,
			password_hash = $2,
			date_birth = $3
		where id = $4;`,
		user.Email,
		user.PasswordHash,
		user.DateBirth,
		user.ID,
	)
	if err != nil {
		repo.logger.Trace(err)
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}
