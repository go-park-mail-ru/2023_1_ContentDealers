package domain

type MovieSelection struct {
	ID     uint64   `json:"id"`
	Movies []*Movie `json:"movies"`
	Title  string   `json:"title"`
}
