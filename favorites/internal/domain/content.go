package domain

import "time"

type FavoritesOptions struct {
	UserID uint64
	Order  string
}

type FavoriteContent struct {
	UserID     uint64
	ContentID  uint64
	DateAdding time.Time
}

type FavoritesContent struct {
	Favorites []FavoriteContent
}
