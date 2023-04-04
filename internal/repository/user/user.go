package user

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const (
	shortFormDate             = "2006-01-02"
	dirAvatars                = "./media/avatars"
	allPerms      os.FileMode = 0777
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (repo *Repository) Add(user domain.User) (domain.User, error) {
	var lastInsertedID uint64
	log.Println(user.Birthday)

	err := repo.DB.QueryRow(
		`insert into users (email, password_hash, birthday, avatar_url) 
        values ($1, $2, $3, $4) 
        returning id`,
		user.Email,
		user.PasswordHash,
		user.Birthday,
		user.AvatarURL,
	).Scan(&lastInsertedID)
	if err != nil {
		// FIXME:
		log.Println("sdfgsdfgdg ", err)
		// TODO: можно ли проверить конкртеную ошибку postgresql (нарушение unique)?
		// https://www.manniwood.com/2016_08_14/pgxfiles_04.html
		// https://stackoverflow.com/questions/70515729/how-to-handle-postgres-query-error-with-pgx-driver-in-golang
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}
	user.ID = lastInsertedID
	return user, nil
}

func (repo *Repository) GetByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.
		QueryRow(`select id, email, password_hash, birthday, avatar_url FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Birthday, &user.AvatarURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (repo *Repository) GetByID(id uint64) (domain.User, error) {
	// TODO: копипаст метода GetByEmail (нужен общий метод для запроса)
	user := domain.User{}
	err := repo.DB.
		QueryRow(`select id, email, password_hash, birthday, avatar_url FROM users WHERE id = $1`, id).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Birthday, &user.AvatarURL)
	if err != nil {
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

func (repo *Repository) deleteAvatar(user domain.User) error {
	var updateDate time.Time

	// 1. удаление из локальной директории
	if user.AvatarURL == "" {
		return nil
	}
	err := repo.DB.
		QueryRow(`select updated_at FROM users WHERE id = $1`, user.ID).
		Scan(&updateDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrUserNotFound
		}
		return err
	}

	filepathForDelete := fmt.Sprintf("%s/%s/%d.jpg", dirAvatars, getDirByDate(updateDate), user.ID)
	if _, err = os.Stat(filepathForDelete); !errors.Is(err, os.ErrNotExist) {
		err = os.Remove(filepathForDelete)
		if err != nil {
			return err
		}
	}

	// 2. удаление урла из БД
	_, err = repo.DB.Exec(
		`update users 
		set avatar_url = null
		where id = $1;`,
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) UpdateAvatar(user domain.User, file io.Reader) (domain.User, error) {
	err := repo.deleteAvatar(user)
	if err != nil {
		return domain.User{}, err
	}

	// TODO: разделитель пути зависит от операционной системы
	dir := fmt.Sprintf("%s/%s", dirAvatars, getDirByDate(time.Now()))

	// если директория существовала, err == nil
	// TODO: хардкод прав на директорию
	err = os.MkdirAll(dir, allPerms)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create folder for avatar")
	}
	filepath := fmt.Sprintf("%s/%d.jpg", dir, user.ID)

	outAvatar, err := os.Create(filepath)
	defer outAvatar.Close()
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create avatar file")
	}
	_, err = io.Copy(outAvatar, file)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to copy avatar file to local directory")
	}

	_, err = repo.DB.Exec(
		`update users 
		set avatar_url = $1
		where id = $2;`,
		filepath,
		user.ID,
	)
	if err != nil {
		return domain.User{}, err
	}
	user.AvatarURL = filepath
	return user, nil
}
