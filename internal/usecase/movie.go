package usecase

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

type MovieRepository interface {
	GetById(id uint64) (domain.Movie, error)
	GetAll() ([]domain.Movie, error)
}
