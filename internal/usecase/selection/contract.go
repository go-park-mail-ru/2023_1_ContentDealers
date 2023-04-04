package selection

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type MovieSelectionRepository interface {
	GetAll() ([]domain.Selection, error)
	GetByID(id uint64) (domain.Selection, error)
}
