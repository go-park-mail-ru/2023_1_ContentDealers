package domain

import "time"

type FavoritesOptions struct {
	UserID   uint64
	SortDate string
}

type FavoriteContent struct {
	UserID     uint64
	ContentID  uint64
	DateAdding time.Time
}
