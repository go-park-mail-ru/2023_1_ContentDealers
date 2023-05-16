package favorites

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainFav "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
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

func (uc *UseCase) Delete(ctx context.Context, favorite domainFav.FavoriteContent) error {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.DeleteFavContent(ctx, favorite)
}

func (uc *UseCase) Add(ctx context.Context, favorite domainFav.FavoriteContent) error {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.AddFavContent(ctx, favorite)
}

func (uc *UseCase) HasFav(ctx context.Context, favorite domainFav.FavoriteContent) (bool, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return false, err
	}
	favorite.UserID = userID
	return uc.gate.HasFavContent(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domainFav.FavoritesOptions) ([]domain.Content, bool, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domain.Content{}, false, err
	}
	options.UserID = userID
	favs, err := uc.gate.GetFavContent(ctx, options)
	if err != nil {
		return []domain.Content{}, false, err
	}

	var contentIDs []uint64
	for _, fav := range favs.Favorites {
		contentIDs = append(contentIDs, fav.ContentID)
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

	return contentSlice, favs.IsLast, nil
}
