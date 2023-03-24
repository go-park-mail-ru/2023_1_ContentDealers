package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type SessionUseCase interface {
	Create(user domain.User) (domain.Session, error)
	Get(sessionID uuid.UUID) (domain.Session, error)
	Delete(sessionID uuid.UUID) error
}
