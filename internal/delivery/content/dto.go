package content

import "time"

type personDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type roleDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type personRolesDTO struct {
	Person personDTO `json:"person"`
	Role   roleDTO   `json:"role"`
}

type genreDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type selectionDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type countryDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type contentDTO struct {
	ID           uint64           `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Rating       float64          `json:"rating"`
	Year         int              `json:"year"`
	IsFree       bool             `json:"is_free"`
	AgeLimit     int              `json:"age_limit"`
	TrailerURL   string           `json:"trailer_url"`
	PreviewURL   string           `json:"preview_url"`
	Type         string           `json:"type"`
	PersonsRoles []personRolesDTO `json:"persons_roles"`
	Genres       []genreDTO       `json:"genres"`
	Selections   []selectionDTO   `json:"selections"`
	Countries    []countryDTO     `json:"countries"`
}

type filmDTO struct {
	ID         uint64     `json:"id"`
	ContentURL string     `json:"content_url"`
	Content    contentDTO `json:"content"`
}

type episodeDTO struct {
	ID          uint64    `json:"id"`
	SeasonNum   uint32    `json:"season_num"`
	EpisodeNum  uint32    `json:"episode_num"`
	ContentURL  string    `json:"content_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
}

type seriesDTO struct {
	ID       uint64       `json:"id"`
	Content  contentDTO   `json:"content"`
	Episodes []episodeDTO `json:"episodes"`
}
