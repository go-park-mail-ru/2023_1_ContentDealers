package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase interface {
	GetByID(ctx context.Context, id uint64) (domain.Person, error)
}
