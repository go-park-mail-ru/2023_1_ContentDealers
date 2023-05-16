package history_views

import (
	"context"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type SessionGateway interface {
	Create(ctx context.Context, user domainUser.User) (domainSession.Session, error)
	Get(ctx context.Context, sessionID string) (domainSession.Session, error)
	Delete(ctx context.Context, sessionID string) error
}

type Gateway interface {
	UpdateProgressView(context.Context, domain.View) error
	GetViewsByUser(context.Context, domain.ViewsOptions) (domain.Views, error)
	HasView(context.Context, domain.View) (domain.HasView, error)
}

type ContentGateway interface {
	GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domainContent.Content, error)
}
