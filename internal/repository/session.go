package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/google/uuid"
	"time"
)

var ErrSessionNotFound = errors.New("session not found")

var _ usecase.SessionRepository = (*SessionInMemoryRepository)(nil)

type SessionInMemoryRepository struct {
	storage map[uuid.UUID]domain.Session
}

func NewSessionInMemoryRepository() SessionInMemoryRepository {
	return SessionInMemoryRepository{storage: map[uuid.UUID]domain.Session{}}
}

func (repo *SessionInMemoryRepository) Add(session domain.Session) error {
	if session.ExpiresAt.Before(time.Now()) {
		return nil
	}
	repo.storage[session.ID] = session
	return nil
}

func (repo *SessionInMemoryRepository) Get(sessionID uuid.UUID) (domain.Session, error) {
	session, ok := repo.storage[sessionID]
	if !ok {
		return session, ErrSessionNotFound
	}
	return session, nil
}

func (repo *SessionInMemoryRepository) Delete(sessionID uuid.UUID) error {
	delete(repo.storage, sessionID)
	return nil
}
