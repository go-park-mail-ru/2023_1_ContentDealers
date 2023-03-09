package domain

import "errors"

var ErrMovieSelectionNotFound = errors.New("movie selection not found")

type MovieSelection struct {
	ID     uint64   `json:"id"`
	Title  string   `json:"title"`
	Movies []*Movie `json:"movies"`
}
