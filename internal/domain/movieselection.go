package domain

type MovieSelection struct {
	ID     uint64   `json:"id"`
	Title  string   `json:"title"`
	Movies []*Movie `json:"movies"`
}
