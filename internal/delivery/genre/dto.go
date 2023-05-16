package genre

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

type genreDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type genreContentDTO struct {
	Genre   genreDTO     `json:"genre"`
	Content []contentDTO `json:"content"`
}
