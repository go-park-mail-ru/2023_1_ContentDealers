package session

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/domain"
)

type SessionUseCase interface {
	Create(ctx context.Context, userID uint64) (domain.Session, error)
	Get(ctx context.Context, sessionID string) (domain.Session, error)
	Delete(ctx context.Context, sessionID string) error
}
