package domain

import "errors"

var ErrSelectionNotFound = errors.New("selection not found")

type SelectionContent struct {
	ID         uint64
	Title      string
	Rating     string
	PreviewURL string
	TrailerURL string
	Type       string // Поле для фронтендеров
	Genre      Genre
}

type Selection struct {
	ID      uint64
	Title   string
	Content []SelectionContent
}
