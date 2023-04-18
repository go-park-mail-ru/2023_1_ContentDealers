package user

type userCreateDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type userUpdateDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userCredentialsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
