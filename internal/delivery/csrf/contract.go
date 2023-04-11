package csrf

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type CSRFUseCase interface {
	Create(domain.Session, int64) (string, error)
	Check(domain.Session, string) (bool, error)
}
