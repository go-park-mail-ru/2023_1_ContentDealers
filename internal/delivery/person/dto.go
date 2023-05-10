package person

type contentDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type roleDTO struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type genreDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type personDTO struct {
	ID             uint64       `json:"id"`
	Name           string       `json:"name"`
	Gender         string       `json:"gender"`
	Growth         *int         `json:"growth"`
	Birthplace     string       `json:"birthplace"`
	AvatarURL      string       `json:"avatar_url"`
	Age            int          `json:"age"`
	ParticipatedIn []contentDTO `json:"participated_in"`
	Roles          []roleDTO    `json:"roles"`
	Genres         []genreDTO   `json:"genres"`
}
