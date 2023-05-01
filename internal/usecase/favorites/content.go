package favorites

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	domainFav "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
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
	return uc.gate.Delete(ctx, favorite)
}

func (uc *UseCase) Add(ctx context.Context, favorite domainFav.FavoriteContent) error {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.Add(ctx, favorite)
}

func (uc *UseCase) HasFav(ctx context.Context, favorite domainFav.FavoriteContent) (bool, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return false, err
	}
	favorite.UserID = userID
	return uc.gate.HasFav(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domainFav.FavoritesOptions) ([]domain.Content, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domain.Content{}, err
	}
	options.UserID = userID
	favs, err := uc.gate.Get(ctx, options)
	if err != nil {
		return []domain.Content{}, err
	}

	var contentIDs []uint64
	for _, fav := range favs {
		contentIDs = append(contentIDs, fav.ContentID)
	}

	contentSliceSorted, err := uc.content.GetContentByContentIDs(ctx, contentIDs)
	if err != nil {
		return []domain.Content{}, err
	}

	// сортировка []domain.Content согласно порядку id-шников в contentIDs

	contentDict := make(map[uint64]domain.Content)
	for _, item := range contentSliceSorted {
		contentDict[item.ID] = item
	}

	contentSlice := make([]domain.Content, len(contentIDs))
	for i, id := range contentIDs {
		contentSlice[i] = contentDict[id]
	}

	return contentSlice, nil
}
