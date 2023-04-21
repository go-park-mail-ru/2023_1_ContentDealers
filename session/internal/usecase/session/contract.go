package session

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/domain"
)

type Repository interface {
	Add(ctx context.Context, session domain.Session) error
	Get(ctx context.Context, id string) (domain.Session, error)
	Delete(ctx context.Context, id string) error
}
