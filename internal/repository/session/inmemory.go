package session

import (
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

var _ contract.SessionRepository = (*InMemoryRepository)(nil)

type InMemoryRepository struct {
	storage map[uuid.UUID]domain.Session
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{storage: map[uuid.UUID]domain.Session{}}
}

func (repo *InMemoryRepository) Add(session domain.Session) error {
	if session.ExpiresAt.Before(time.Now()) {
		return nil
	}
	repo.storage[session.ID] = session
	return nil
}

func (repo *InMemoryRepository) Get(sessionID uuid.UUID) (domain.Session, error) {
	session, ok := repo.storage[sessionID]
	if !ok {
		return session, domain.ErrSessionNotFound
	}
	return session, nil
}

func (repo *InMemoryRepository) Delete(sessionID uuid.UUID) error {
	delete(repo.storage, sessionID)
	return nil
}
