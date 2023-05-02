package domain

type Genre struct {
	ID   uint64
	Name string
}

type GenreContent struct {
	Genre   Genre
	Content []Content
}
