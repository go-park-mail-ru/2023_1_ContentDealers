package setup

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type SessionUseCase interface {
	Create(ctx context.Context, user domain.User) (domain.Session, error)
	Get(ctx context.Context, sessionID uuid.UUID) (domain.Session, error)
	Delete(ctx context.Context, sessionID uuid.UUID) error
}
