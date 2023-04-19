package user

import (
	"context"
	"io"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/google/uuid"
)

type UserUseCase interface {
	Register(ctx context.Context, user domain2.User) (domain2.User, error)
	Auth(ctx context.Context, user domain2.User) (domain2.User, error)
	GetByID(ctx context.Context, id uint64) (domain2.User, error)
	Update(ctx context.Context, user domain2.User) error
	UpdateAvatar(context.Context, domain2.User, io.Reader) (domain2.User, error)
	DeleteAvatar(context.Context, domain2.User) error
}

type SessionUseCase interface {
	Create(user domain2.User) (domain2.Session, error)
	Get(sessionID uuid.UUID) (domain2.Session, error)
	Delete(sessionID uuid.UUID) error
}
