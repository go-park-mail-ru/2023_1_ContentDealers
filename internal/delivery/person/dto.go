package person

import "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"

var a = domain.Person{
	ID:             0,
	Name:           "",
	Gender:         "",
	Growth:         nil,
	Birthplace:     nil,
	AvatarURL:      "",
	Age:            0,
	ParticipatedIn: nil,
	Roles:          nil,
	Genres:         nil,
}

type contentDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type roleDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type genreDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type personDTO struct {
	ID             uint64       `json:"id"`
	Name           string       `json:"name"`
	Gender         string       `json:"gender"`
	Growth         *int         `json:"growth"`
	Birthplace     string       `json:"birthplace"`
	AvatarURL      string       `json:"avatar_url"`
	ParticipatedIn []contentDTO `json:"participated_in"`
	Roles          []roleDTO    `json:"roles"`
	Genres         []genreDTO   `json:"genres"`
}
