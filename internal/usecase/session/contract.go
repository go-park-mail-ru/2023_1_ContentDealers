package session

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Add(session domain.Session) error
	Get(id uuid.UUID) (domain.Session, error)
	Delete(id uuid.UUID) error
}
