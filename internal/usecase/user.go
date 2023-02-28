package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

var ErrWrongCredentials = errors.New("wrong credentials")

type UserRepository interface {
	Add(user domain.UserCredentials) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetById(id uint64) (domain.User, error)
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) RegisterUser(credentials domain.UserCredentials) (domain.User, error) {
	return uc.repo.Add(credentials)
}

func (uc *UserUseCase) AuthUser(credentials domain.UserCredentials) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(credentials.Email)
	if err != nil {
		return domain.User{}, err
	}
	if realUser.Password != credentials.Password {
		return domain.User{}, ErrWrongCredentials
	}
	return realUser, nil
}

func (uc *UserUseCase) GetById(id uint64) (domain.User, error) {
	return uc.repo.GetById(id)
}
