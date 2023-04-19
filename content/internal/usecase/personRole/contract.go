package personRole

import (
	"context"

	domain2 "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
)

type PersonRepository interface {
	GetByContentID(ctx context.Context, ContentID uint64) ([]domain2.Person, error)
}

type RoleRepository interface {
	GetByContentID(ct context.Context, ContentID uint64) (map[uint64]domain2.Role, error)
}
