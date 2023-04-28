package favorites

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	gate    Gateway
	session SessionUseCase
	content ContentUseCase
	logger  logging.Logger
}

func NewUseCase(gate Gateway, session SessionUseCase, content ContentUseCase, logger logging.Logger) *UseCase {
	return &UseCase{gate: gate, session: session, content: content, logger: logger}
}

func (uc *UseCase) GetUserIDByContext(ctx context.Context) (uint64, error) {
	session, ok := ctx.Value("session").(domain.Session)
	if !ok {
		return 0, fmt.Errorf("session not found")
	}
	return session.UserID, nil
}

func (uc *UseCase) Delete(ctx context.Context, favorite domain.FavoriteContent) error {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.Delete(ctx, favorite)
}

func (uc *UseCase) Add(ctx context.Context, favorite domain.FavoriteContent) error {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	favorite.UserID = userID
	return uc.gate.Add(ctx, favorite)
}

func (uc *UseCase) Get(ctx context.Context, options domain.FavoritesOptions) ([]domain.FavoriteContent, error) {
	userID, err := uc.GetUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domain.FavoriteContent{}, err
	}
	options.UserID = userID
	return uc.gate.Get(ctx, options)
}
