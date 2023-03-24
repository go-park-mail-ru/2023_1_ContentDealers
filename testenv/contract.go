package testenv

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type MovieRepository interface {
	GetByID(id uint64) (domain.Movie, error)
	GetAll() ([]domain.Movie, error)
}

type MovieSelectionRepository interface {
	GetAll() ([]domain.MovieSelection, error)
	GetByID(id uint64) (domain.MovieSelection, error)
}

type MovieSelectionUseCase interface {
	GetAll() ([]domain.MovieSelection, error)
	GetByID(id uint64) (domain.MovieSelection, error)
}

type SessionRepository interface {
	Add(session domain.Session) error
	Get(id uuid.UUID) (domain.Session, error)
	Delete(id uuid.UUID) error
}

type SessionUseCase interface {
	Create(user domain.User) (domain.Session, error)
	Get(sessionID uuid.UUID) (domain.Session, error)
	Delete(sessionID uuid.UUID) error
}

type UserRepository interface {
	Add(user domain.UserCredentials) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetByID(id uint64) (domain.User, error)
}

type UserUseCase interface {
	Register(credentials domain.UserCredentials) (domain.User, error)
	Auth(credentials domain.UserCredentials) (domain.User, error)
	GetByID(id uint64) (domain.User, error)
}
