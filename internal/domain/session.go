package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrSessionNotFound = errors.New("session not found")
var ErrSessionInvalid = errors.New("fail to cast session")

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
