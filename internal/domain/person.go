package domain

type Person struct {
	ID             uint64
	Name           string
	Gender         rune
	Growth         *int
	Birthplace     *string
	AvatarURL      string
	Age            int
	ParticipatedIn []Content
	Roles          []Role
	Genres         []Genre
}
