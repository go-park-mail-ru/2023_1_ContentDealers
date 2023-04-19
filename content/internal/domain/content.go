package domain

type PersonRoles struct {
	Person Person
	Role   Role
}

type Content struct {
	ID           uint64        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Rating       float64       `json:"rating"`
	Year         int           `json:"year"`
	IsFree       bool          `json:"is_free"`
	AgeLimit     int           `json:"age_limit"`
	TrailerURL   string        `json:"trailer_url"`
	PreviewURL   string        `json:"preview_url"`
	Type         string        `json:"type"`
	PersonsRoles []PersonRoles `json:"persons_roles"`
	Genres       []Genre       `json:"genres"`
	Selections   []Selection   `json:"selections"`
	Countries    []Country     `json:"countries"`
}
