package domain

type User struct {
	ID uint64
	UserCredentials
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
