package session

import (
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/google/uuid"
)

const SessionTimeout = time.Hour * 12

type Session struct {
	repo   Repository
	logger logging.Logger
}

func NewSession(repo Repository, logger logging.Logger) *Session {
	return &Session{repo: repo, logger: logger}
}

func (uc *Session) Create(user domain.User) (domain.Session, error) {
	newSession := domain.NewSession(user.ID, time.Now().Add(SessionTimeout))
	err := uc.repo.Add(newSession)
	return newSession, err
}

func (uc *Session) Get(sessionID uuid.UUID) (domain.Session, error) {
	return uc.repo.Get(sessionID)
}

func (uc *Session) Delete(sessionID uuid.UUID) error {
	return uc.repo.Delete(sessionID)
}
