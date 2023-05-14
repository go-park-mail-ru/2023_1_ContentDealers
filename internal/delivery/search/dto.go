package search

type personDTO struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Growth     *int   `json:"growth"`
	Birthplace string `json:"birthplace"`
	AvatarURL  string `json:"avatar_url"`
	Age        int    `json:"age"`
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

type searchDTO struct {
	Content []contentDTO `json:"content"`
	Persons []personDTO  `json:"persons"`
}
