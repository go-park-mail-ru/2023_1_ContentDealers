package domain

type PersonRoles struct {
	Person Person
	Role   Role
}

type Content struct {
	ID           uint64
	Title        string
	Description  string
	Rating       float64
	SumRatings   float64
	CountRatings uint64
	Year         int
	IsFree       bool
	AgeLimit     int
	TrailerURL   string
	PreviewURL   string
	Type         string
	PersonsRoles []PersonRoles
	Genres       []Genre
	Selections   []Selection
	Countries    []Country
}
