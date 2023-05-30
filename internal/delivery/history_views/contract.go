package history_views

import (
	"context"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	domainView "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type ViewsUseCase interface {
	UpdateProgressView(context.Context, domain.View) error
	GetViewsByUser(context.Context, domain.ViewsOptions) ([]domainContent.Content, []domainView.View, bool, error)
	HasView(context.Context, domain.View) (domain.HasView, error)
}
