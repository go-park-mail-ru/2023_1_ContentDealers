package user

import (
	"errors"
	"sync"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type InMemoryRepository struct {
	mu        sync.RWMutex
	storage   []domain.User
	currentID uint64
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{}
}

func (repo *InMemoryRepository) Add(user domain.UserCredentials) (domain.User, error) {
	if _, err := repo.GetByEmail(user.Email); !errors.Is(err, domain.ErrUserNotFound) {
		return domain.User{}, domain.ErrUserAlreadyExists
	}
	repo.mu.Lock()
	defer repo.mu.Unlock()
	toAdd := domain.User{
		ID:              repo.currentID,
		UserCredentials: user,
	}
	repo.currentID++
	repo.storage = append(repo.storage, toAdd)
	return toAdd, nil
}

func (repo *InMemoryRepository) GetAll() ([]domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	return repo.storage, nil
}

func (repo *InMemoryRepository) GetByEmail(email string) (domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	for _, user := range repo.storage {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}

func (repo *InMemoryRepository) GetByID(id uint64) (domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	for _, user := range repo.storage {
		if user.ID == id {
			return user, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}
