package setup

import (
	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/google/uuid"
)

type SessionUseCase interface {
	Create(user domain2.User) (domain2.Session, error)
	Get(sessionID uuid.UUID) (domain2.Session, error)
	Delete(sessionID uuid.UUID) error
}
