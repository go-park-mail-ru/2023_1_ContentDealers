package user

import (
	"context"
	"io"

	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type UserGateway interface {
	Register(ctx context.Context, user domainUser.User) (domainUser.User, error)
	Auth(ctx context.Context, user domainUser.User) (domainUser.User, error)
	GetByID(ctx context.Context, id uint64) (domainUser.User, error)
	Update(ctx context.Context, user domainUser.User) error
	UpdateAvatar(context.Context, domainUser.User, io.Reader) (domainUser.User, error)
	DeleteAvatar(context.Context, domainUser.User) error
}

type SessionUseCase interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}
