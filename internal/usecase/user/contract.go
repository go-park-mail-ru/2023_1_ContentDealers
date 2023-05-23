package user

import (
	"context"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type ContentGateway interface {
	GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domainContent.Content, error)
}

type SessionGateway interface {
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
}

type UserGateway interface {
	GetByID(ctx context.Context, id uint64) (domainUser.User, error)
}
