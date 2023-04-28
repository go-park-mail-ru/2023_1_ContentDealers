package film

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Film struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewFilm(contentGateway ContentGateway, logger logging.Logger) *Film {
	return &Film{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *Film) GetByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	return uc.contentGateway.GetFilmByContentID(ctx, ContentID)
}
