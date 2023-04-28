package favorites

type FavoriteContentDTO struct {
	UserID    uint64 `json:"user_id"`
	ContentID uint64 `json:"content_id"`
}
