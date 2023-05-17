package domain

import "time"

type RatingsOptions struct {
	UserID    uint64
	ContentID uint64
	Rating    float32
	SortDate  string
	Limit     uint32
	Offset    uint32
}

type Rating struct {
	UserID     uint64
	ContentID  uint64
	Rating     float32
	DateAdding time.Time
}

type HasRating struct {
	Rating    Rating
	HasRating bool
}

type Ratings struct {
	IsLast  bool
	Ratings []Rating
}
