package favorites

import (
	"context"
	"fmt"

	domainFav "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
)

type UseCase struct {
	gate    Gateway
	session SessionGateway
	content ContentUseCase
	logger  logging.Logger
}

func NewUseCase(gate Gateway, session SessionGateway, content ContentUseCase, logger logging.Logger) *UseCase {
	return &UseCase{gate: gate, session: session, content: content, logger: logger}
}

func (uc *UseCase) GetUserIDByContext(ctx context.Context) (uint64, error) {
	session, ok := ctx.Value("session").(domainSession.Session)
	if !ok {
		return 0, fmt.Errorf("session not found")
	}
	return session.UserID, nil
}

func (uc *UseCase) Delete(ctx context.Context, favorite domainFav.FavoriteContent) error {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.Delete(ctx, favorite)
}

func (uc *UseCase) Add(ctx context.Context, favorite domainFav.FavoriteContent) error {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.Add(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domainFav.FavoritesOptions) ([]domainFav.FavoriteContent, error) {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domainFav.FavoriteContent{}, err
	}
	options.UserID = userID
	return uc.gate.Get(ctx, options)
}