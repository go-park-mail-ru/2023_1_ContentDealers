package domain

type PersonContent struct {
	ContentID uint64
	Title     string
}

type Person struct {
	ID             uint64
	Name           string
	Gender         rune
	Growth         *int
	Birthplace     *string
	AvatarURL      string
	Age            int
	ParticipatedIn []PersonContent
	Roles          []Role
	Genres         []Genre
}
