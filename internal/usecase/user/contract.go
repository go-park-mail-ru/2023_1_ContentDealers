package user

import (
	"io"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Repository interface {
	Add(user domain.User) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetByID(id uint64) (domain.User, error)
	UpdateAvatar(domain.User, io.Reader) (domain.User, error)
}
