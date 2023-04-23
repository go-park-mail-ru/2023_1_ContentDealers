package selection

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Selection struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewSelection(contentGateway ContentGateway, logger logging.Logger) *Selection {
	return &Selection{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *Selection) GetAll(ctx context.Context, limit, offset uint32) ([]domain.Selection, error) {
	return uc.contentGateway.GetAllSelections(ctx, limit, offset)
}

func (uc *Selection) GetByID(ctx context.Context, id uint64) (domain.Selection, error) {
	return uc.contentGateway.GetSelectionByID(ctx, id)
}
