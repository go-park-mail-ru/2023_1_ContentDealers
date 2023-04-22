package session

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/domain"
)

type Session struct {
	repo      Repository
	logger    logging.Logger
	expiresAt time.Duration
}

func NewSession(repo Repository, logger logging.Logger, expiresAt int) *Session {
	return &Session{repo: repo, logger: logger, expiresAt: time.Duration(expiresAt) * time.Second}
}

func (uc *Session) Create(ctx context.Context, userID uint64) (domain.Session, error) {
	newSession := domain.NewSession(userID, time.Now().Add(uc.expiresAt))
	err := uc.repo.Add(ctx, newSession)
	return newSession, err
}

func (uc *Session) Get(ctx context.Context, sessionID string) (domain.Session, error) {
	return uc.repo.Get(ctx, sessionID)
}

func (uc *Session) Delete(ctx context.Context, sessionID string) error {
	return uc.repo.Delete(ctx, sessionID)
}
