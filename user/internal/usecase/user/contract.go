package user

import (
	"context"
	"io"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type Repository interface {
	Add(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByID(ctx context.Context, id uint64) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	UpdateAvatar(context.Context, domain.User, io.Reader) (domain.User, error)
	DeleteAvatar(ctx context.Context, user domain.User) error
	Subscribe(ctx context.Context, user domain.User) error
}
