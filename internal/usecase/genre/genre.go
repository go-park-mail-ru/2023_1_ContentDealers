package genre

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewUseCase(contentGateway ContentGateway, logger logging.Logger) *UseCase {
	return &UseCase{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *UseCase) GetGenreContent(ctx context.Context, filter domain.ContentFilter) (domain.GenreContent, error) {
	return uc.contentGateway.GetGenreContent(ctx, filter)
}

func (uc *UseCase) GetAll(ctx context.Context) ([]domain.Genre, error) {
	return uc.contentGateway.GetAllGenres(ctx)
}
