package movieselection

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type MovieSelectionRepository interface {
	GetAll() ([]domain.MovieSelection, error)
	GetByID(id uint64) (domain.MovieSelection, error)
}
