package user

import (
	"io"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type UserUseCase interface {
	Register(user domain.User) (domain.User, error)
	Auth(user domain.User) (domain.User, error)
	GetByID(id uint64) (domain.User, error)
	UpdateAvatar(domain.User, io.Reader) (domain.User, error)
}

type SessionUseCase interface {
	Create(user domain.User) (domain.Session, error)
	Get(sessionID uuid.UUID) (domain.Session, error)
	Delete(sessionID uuid.UUID) error
}
