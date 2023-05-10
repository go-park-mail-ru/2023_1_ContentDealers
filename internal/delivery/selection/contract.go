package selection

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase interface {
	GetAll(ctx context.Context, limit, offset uint32) ([]domain.Selection, error)
	GetByID(ctx context.Context, id uint64) (domain.Selection, error)
}
