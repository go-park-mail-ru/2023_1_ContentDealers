package selection

type contentDTO struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Year        int     `json:"year"`
	IsFree      bool    `json:"is_free"`
	AgeLimit    int     `json:"age_limit"`
	TrailerURL  string  `json:"trailer_url"`
	PreviewURL  string  `json:"preview_url"`
	Type        string  `json:"type"`
}

type selectionDTO struct {
	ID      uint64       `json:"id"`
	Title   string       `json:"title"`
	Content []contentDTO `json:"content"`
}
