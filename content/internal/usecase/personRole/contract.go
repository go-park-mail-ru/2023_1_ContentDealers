package personRole

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type PersonRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain.Person, error)
}

type RoleRepository interface {
	GetByContentID(ct context.Context, ContentID uint64) (map[uint64]domain.Role, error)
}
