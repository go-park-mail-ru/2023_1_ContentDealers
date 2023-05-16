package domain

import "time"

type Episode struct {
	ID          uint64
	SeasonNum   uint32
	EpisodeNum  uint32
	ContentURL  string
	PreviewURL  string
	Title       *string
	ReleaseDate *time.Time
}
