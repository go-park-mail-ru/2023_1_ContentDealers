package user_action

import "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"

type FavoritesContentDTO struct {
	Favorites []domain.FavoriteContent
	IsLast    bool
}
