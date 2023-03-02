package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
)

func Content(movieRepo *repository.MovieInMemoryRepository,
	movieSelectionRepo *repository.MovieSelectionInMemoryRepository) {
	movies := []domain.Movie{
		{ID: 0, Title: "Легенда", Description: "Криминальная драма о двух братьях-гангстерах из Лондона"},
		{ID: 1, Title: "Операция Ы", Description: "Новые приключения Шурик и знаменитой тройки"},
	}
	selections := []domain.MovieSelection{
		{
			ID:    0,
			Title: "Filmium рекомендует",
			Movies: []*domain.Movie{
				&movies[0],
			},
		},
		{
			ID:    1,
			Title: "Советская классика",
			Movies: []*domain.Movie{
				&movies[1],
			},
		},
	}
	movieRepo.Add(movies)
	movieSelectionRepo.Add(selections)
}
