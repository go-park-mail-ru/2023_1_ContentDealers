package person

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type UseCase interface {
	GetByID(id uint64) (domain.Person, error)
}
