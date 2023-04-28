package setup

import (
	"context"

	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type SessionUseCase interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}
