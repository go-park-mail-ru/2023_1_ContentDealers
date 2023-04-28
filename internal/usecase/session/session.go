package session

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

const SessionTimeout = time.Hour * 12

type Session struct {
	gateway Gateway
	logger  logging.Logger
}

func NewSession(gateway Gateway, logger logging.Logger) *Session {
	return &Session{gateway: gateway, logger: logger}
}

func (uc *Session) Create(ctx context.Context, user domainUser.User) (domainSession.Session, error) {
	return uc.gateway.Create(ctx, user)
}

func (uc *Session) Get(ctx context.Context, sessionID string) (domainSession.Session, error) {
	return uc.gateway.Get(ctx, sessionID)
}

func (uc *Session) Delete(ctx context.Context, sessionID string) error {
	return uc.gateway.Delete(ctx, sessionID)
}
