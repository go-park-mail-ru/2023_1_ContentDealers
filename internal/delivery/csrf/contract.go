package csrf

import (
	"context"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
)

type UseCase interface {
	Create(context.Context, domain.Session, int64) (string, error)
	Check(context.Context, domain.Session, string) (bool, error)
}
