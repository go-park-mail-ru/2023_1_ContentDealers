package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// Проверка соответствию интерфейса
var _ usecase.UserRepository = (*UserInMemoryRepository)(nil)

type UserInMemoryRepository struct {
	storage   []domain.User
	currentID uint64
}

func NewUserInMemoryRepository() UserInMemoryRepository {
	return UserInMemoryRepository{}
}

func (repo *UserInMemoryRepository) Add(user domain.UserCredentials) (domain.User, error) {
	if _, err := repo.GetByEmail(user.Email); err != ErrUserNotFound {
		return domain.User{}, ErrUserAlreadyExists
	}
	toAdd := domain.User{
		ID:              repo.currentID,
		UserCredentials: user,
	}
	repo.currentID++
	repo.storage = append(repo.storage, toAdd)
	return toAdd, nil
}

func (repo *UserInMemoryRepository) GetAll() ([]domain.User, error) {
	return repo.storage, nil
}

func (repo *UserInMemoryRepository) GetByEmail(email string) (domain.User, error) {
	for _, user := range repo.storage {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, ErrUserNotFound
}

func (repo *UserInMemoryRepository) GetById(id uint64) (domain.User, error) {
	for _, user := range repo.storage {
		if user.ID == id {
			return user, nil
		}
	}
	return domain.User{}, ErrUserNotFound
}
