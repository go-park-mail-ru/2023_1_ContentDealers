package history_views

type viewDTO struct {
	UserID    uint64 `json:"user_id"`
	ContentID uint64 `json:"content_id"`
	StopView  string `json:"stop_view"`
	Duration  string `json:"duration"`
}

type contentViewDTO struct {
	Content contentDTO `json:"content"`
	View    viewDTO    `json:"view"`
}

type contentDTO struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	SumRating   float64 `json:"sum_rating"`
	CountRating uint64  `json:"count_rating"`
	Year        int     `json:"year"`
	IsFree      bool    `json:"is_free"`
	AgeLimit    int     `json:"age_limit"`
	TrailerURL  string  `json:"trailer_url"`
	PreviewURL  string  `json:"preview_url"`
	Type        string  `json:"type"`
}
