package user

import (
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository struct {
	mu        sync.RWMutex
	storage   []domain.User
	currentID uint64
	DB        *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

// TODO: кажется, что DTO здесь будет очень лишним
func (repo *Repository) Add(user domain.UserCredentials) (domain.User, error) {
	var lastInsertId uint64
	err := repo.DB.QueryRow(
		`insert into users(email, password_hash) 
        values ($1, $2) 
        returning id`,
		user.Email,
		user.Password,
	).Scan(&lastInsertId)
	if err != nil {
		// TODO: можно ли проверить конкретную ошибку postgresql (нарушение unique)?
		// https://www.manniwood.com/2016_08_14/pgxfiles_04.html
		// https://stackoverflow.com/questions/70515729/how-to-handle-postgres-query-error-with-pgx-driver-in-golang
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, domain.ErrUserAlreadyExists
		}
		return domain.User{}, err
	}
	toAdd := domain.User{
		ID:              lastInsertId,
		UserCredentials: user,
	}
	return toAdd, nil
}

func (repo *Repository) GetByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.
		QueryRow(`select id, email, password_hash FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (repo *Repository) GetByID(id uint64) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.
		QueryRow(`select id, email FROM users WHERE id = $1`, id).
		Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}
