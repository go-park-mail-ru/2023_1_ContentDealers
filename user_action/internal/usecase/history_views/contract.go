package history_views

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type Repository interface {
	UpdateProgressView(context.Context, domain.View) error
	GetViewsByUser(context.Context, domain.ViewsOptions) (domain.Views, error)
	HasView(context.Context, domain.View) (domain.HasView, error)
}
