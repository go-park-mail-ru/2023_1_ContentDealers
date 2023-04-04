package domain

import "errors"

var ErrFilmNotFound = errors.New("film not found")

type Film struct {
	ID         uint64
	ContentURL string
	Content
}
