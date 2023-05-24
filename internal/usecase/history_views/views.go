package history_views

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainView "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type UseCase struct {
	gate    Gateway
	session SessionGateway
	logger  logging.Logger
	content ContentGateway
}

func NewUseCase(gate Gateway, session SessionGateway, content ContentGateway, logger logging.Logger) *UseCase {
	return &UseCase{gate: gate, session: session, content: content, logger: logger}
}

func (uc *UseCase) getUserIDByContext(ctx context.Context) (uint64, error) {
	session, ok := ctx.Value("session").(domainSession.Session)
	if !ok {
		return 0, fmt.Errorf("session not found")
	}
	return session.UserID, nil
}

func (uc *UseCase) UpdateProgressView(ctx context.Context, view domainView.View) error {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	view.UserID = userID
	err = uc.gate.UpdateProgressView(ctx, view)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (uc *UseCase) HasView(ctx context.Context, view domainView.View) (domainView.HasView, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return domainView.HasView{}, err
	}
	view.UserID = userID
	return uc.gate.HasView(ctx, view)
}

func (uc *UseCase) GetViewsByUser(ctx context.Context, options domainView.ViewsOptions) ([]domainContent.Content, bool, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domain.Content{}, false, err
	}
	options.UserID = userID
	views, err := uc.gate.GetViewsByUser(ctx, options)
	if err != nil {
		return []domain.Content{}, false, err
	}

	contentIDs := make([]uint64, 0, len(views.Views))
	for _, view := range views.Views {
		contentIDs = append(contentIDs, view.ContentID)
	}

	contentSliceSorted, err := uc.content.GetContentByContentIDs(ctx, contentIDs)
	if err != nil {
		return []domain.Content{}, false, err
	}

	// сортировка contentSliceSorted согласно порядку id-шников в contentIDs

	contentDict := make(map[uint64]domain.Content)
	for _, item := range contentSliceSorted {
		contentDict[item.ID] = item
	}

	contentSlice := make([]domain.Content, len(contentIDs))
	for i, id := range contentIDs {
		contentSlice[i] = contentDict[id]
	}

	return contentSlice, views.IsLast, nil
}
