package usecase

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
	"time"
)

type SessionRepository interface {
	Add(session domain.Session) error
	Get(id uuid.UUID) (domain.Session, error)
	Delete(id uuid.UUID) error
}

type SessionUseCase struct {
	repo SessionRepository
}

func NewSessionUseCase(repo SessionRepository) *SessionUseCase {
	return &SessionUseCase{repo: repo}
}

func (uc *SessionUseCase) CreateSession(user domain.User) (domain.Session, error) {
	newSession := domain.NewSession(user.ID, time.Now().Add(time.Hour*12))
	err := uc.repo.Add(newSession)
	return newSession, err
}

func (uc *SessionUseCase) GetSession(sessionID uuid.UUID) (domain.Session, error) {
	return uc.repo.Get(sessionID)
}

func (uc *SessionUseCase) DeleteSession(sessionID uuid.UUID) error {
	return uc.repo.Delete(sessionID)
}
