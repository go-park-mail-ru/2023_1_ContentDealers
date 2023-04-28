package session

import (
	"context"

	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

type Gateway interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, id string) (domainSession.Session, error)
	Delete(ctx context.Context, id string) error
}
