package domain

type ContentPerson struct {
	ID   uint64
	Name string
	Role Role
}

type Content struct {
	ID           uint64
	Title        string
	Description  string
	Rating       float64
	Year         int
	IsFree       bool
	AgeLimit     int
	TrailerURL   string
	PreviewURL   string
	Type         string
	PersonsRoles []ContentPerson
	Genres       []Genre
	Selections   []Selection
	Countries    []Country
}
