package session

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Add(ctx context.Context, session domain.Session) error
	Get(ctx context.Context, id uuid.UUID) (domain.Session, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
