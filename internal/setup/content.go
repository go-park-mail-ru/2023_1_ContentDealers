package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
)

var Movies = []domain.Movie{
	{ID: 0, PreviewURL: "previews/mad-max.jpg", Title: "Mad Max", Description: "Mad Max описание фильма"},
	{ID: 1, PreviewURL: "previews/back-to-the-future.jpg", Title: "Back to the future", Description: "Back to the future описание фильма"},
	{ID: 2, PreviewURL: "previews/king-kong.jpg", Title: "King Kong", Description: "King Kong описание фильма"},
	{ID: 3, PreviewURL: "previews/terminator.jpg", Title: "Terminator", Description: "Terminator описание фильма"},
	{ID: 4, PreviewURL: "previews/godzilla.jpg", Title: "Godzilla", Description: "Godzilla описание фильма"},
	{ID: 5, PreviewURL: "previews/007.jpg", Title: "007", Description: "007 описание фильма"},
	{ID: 6, PreviewURL: "previews/black-panther.jpg", Title: "Back Panther", Description: "Back Panther описание фильма"},
	{ID: 7, PreviewURL: "previews/captain-america.jpg", Title: "Capitan America", Description: "Capitan America описание фильма"},
	{ID: 8, PreviewURL: "previews/pacific-rim.jpg", Title: "Pacific Rim", Description: "Pacific Rim описание фильма"},
	{ID: 9, PreviewURL: "previews/interstellar.jpg", Title: "Interstellar", Description: "Interstellar описание фильма"},
	{ID: 10, PreviewURL: "previews/face.jpg", Title: "Face", Description: "Face описание фильма"},
	{ID: 11, PreviewURL: "previews/thor.jpg", Title: "Thor", Description: "Thor описание фильма"},
	{ID: 12, PreviewURL: "previews/dune.jpg", Title: "Dune", Description: "Dune описание фильма"},
	{ID: 13, PreviewURL: "previews/avatar.jpg", Title: "Avatar", Description: "Avatar описание фильма"},
	{ID: 14, PreviewURL: "previews/star-wars.jpg", Title: "Star Wars", Description: "Star Wars описание фильма"},
	{ID: 15, PreviewURL: "previews/venom.jpg", Title: "Vanom", Description: "Vanom описание фильма"},
}

var MovieSelections = []domain.MovieSelection{
	{
		ID:    0,
		Title: "Filmium рекомендует",
		Movies: []*domain.Movie{
			&Movies[1],
			&Movies[2],
			&Movies[5],
			&Movies[6],
			&Movies[7],
			&Movies[13],
		},
	},
	{
		ID:    1,
		Title: "Лучшие боевики",
		Movies: []*domain.Movie{
			&Movies[0],
			&Movies[3],
			&Movies[4],
			&Movies[5],
			&Movies[8],
		},
	},
	{
		ID:    3,
		Title: "Фэнтези",
		Movies: []*domain.Movie{
			&Movies[9],
			&Movies[12],
			&Movies[13],
			&Movies[14],
		},
	},
}

func Content(movieRepo *repository.MovieInMemoryRepository,
	movieSelectionRepo *repository.MovieSelectionInMemoryRepository) {
	for _, movie := range Movies {
		movieRepo.Add(movie)
	}
	for _, selection := range MovieSelections {
		movieSelectionRepo.Add(selection)
	}
}
