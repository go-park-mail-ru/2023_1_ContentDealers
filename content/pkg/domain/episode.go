package domain

type Episode struct {
	ID         uint64
	SeriesID   uint64
	SeasonNum  uint64
	ContentURL string
	Title      string
}
