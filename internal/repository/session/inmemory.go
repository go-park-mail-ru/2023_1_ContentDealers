package session

import (
	"sync"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type InMemoryRepository struct {
	mu      sync.RWMutex
	storage map[uuid.UUID]domain.Session
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{storage: map[uuid.UUID]domain.Session{}}
}

func (repo *InMemoryRepository) Add(session domain.Session) error {
	if session.ExpiresAt.Before(time.Now()) {
		return nil
	}
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.storage[session.ID] = session
	return nil
}

func (repo *InMemoryRepository) Get(sessionID uuid.UUID) (domain.Session, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	session, ok := repo.storage[sessionID]
	if !ok {
		return session, domain.ErrSessionNotFound
	}
	return session, nil
}

func (repo *InMemoryRepository) Delete(sessionID uuid.UUID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	delete(repo.storage, sessionID)
	return nil
}
