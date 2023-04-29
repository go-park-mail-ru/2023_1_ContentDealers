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
	film    FilmUseCase
}

func NewUseCase(gate Gateway, session SessionGateway, film FilmUseCase, logger logging.Logger) *UseCase {
	return &UseCase{gate: gate, session: session, film: film, logger: logger}
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

func (uc *UseCase) HasFav(ctx context.Context, favorite domainFav.FavoriteContent) (bool, error) {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return false, err
	}
	favorite.UserID = userID
	return uc.gate.HasFav(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domainFav.FavoritesOptions) ([]domain.Content, error) {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domain.Content{}, err
	}
	options.UserID = userID
	favs, err := uc.gate.Get(ctx, options)
	if err != nil {
		return []domain.Content{}, err
	}

	var content []domain.Content

	for _, fav := range favs {
		film, err := uc.film.GetByContentID(ctx, fav.ContentID)
		if err != nil {
			return []domain.Content{}, err
		}
		content = append(content, film.Content)
	}

	return content, nil
}
