package favorites

type FavoriteContentDTO struct {
	UserID    uint64 `json:"user_id"`
	ContentID uint64 `json:"content_id"`
}

type contentDTO struct {
	ID           uint64  `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Rating       float64 `json:"rating"`
	SumRatings   float64 `json:"sum_ratings"`
	CountRatings uint64  `json:"count_ratings"`
	Year         int     `json:"year"`
	IsFree       bool    `json:"is_free"`
	AgeLimit     int     `json:"age_limit"`
	TrailerURL   string  `json:"trailer_url"`
	PreviewURL   string  `json:"preview_url"`
	Type         string  `json:"type"`
}
