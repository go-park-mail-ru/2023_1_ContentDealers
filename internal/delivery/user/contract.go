package user

import (
	"context"
	"io"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type UserUseCase interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Auth(ctx context.Context, user domain.User) (domain.User, error)
	GetByID(ctx context.Context, id uint64) (domain.User, error)
	UpdateAvatar(context.Context, domain.User, io.Reader) (domain.User, error)
	DeleteAvatar(context.Context, domain.User) error
}

type SessionUseCase interface {
	Create(user domain.User) (domain.Session, error)
	Get(sessionID uuid.UUID) (domain.Session, error)
	Delete(sessionID uuid.UUID) error
}
