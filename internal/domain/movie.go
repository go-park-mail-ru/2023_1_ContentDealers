package domain

import "errors"

var ErrMovieNotFound = errors.New("movie not found")

type Movie struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PreviewURL  string `json:"preview_url"`
}
