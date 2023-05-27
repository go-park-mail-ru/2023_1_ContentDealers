package user

type userDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type userPasswordDTO struct {
	Password string `json:"password"`
}
