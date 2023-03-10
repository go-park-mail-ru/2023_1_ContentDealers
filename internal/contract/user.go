package contract

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

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
