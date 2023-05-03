package domain

import "time"

type FavoritesOptions struct {
	UserID   uint64
	SortDate string
	Limit    uint32
	Offset   uint32
}

type FavoriteContent struct {
	UserID     uint64
	ContentID  uint64
	DateAdding time.Time
}

type FavoritesContent struct {
	IsLast    bool
	Favorites []FavoriteContent
}
