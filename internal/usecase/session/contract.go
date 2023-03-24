package session

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Add(session domain.Session) error
	Get(id uuid.UUID) (domain.Session, error)
	Delete(id uuid.UUID) error
}
