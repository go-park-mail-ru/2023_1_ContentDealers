package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID
	UserID    uint64
	ExpiresAt time.Time
}

func NewSession(UserID uint64, ExpiresAt time.Time) Session {
	return Session{
		ID:        uuid.New(),
		UserID:    UserID,
		ExpiresAt: ExpiresAt,
	}
}
