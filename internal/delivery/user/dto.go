package user

type userCreateDTO struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
	Birthday  string `json:"date_birth"`
}

type userUpdateDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"date_birth"`
}

type userCredentialsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
