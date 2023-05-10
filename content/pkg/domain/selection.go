package domain

type Selection struct {
	ID      uint64    `json:"id"`
	Title   string    `json:"title"`
	Content []Content `json:"content"`
}
