package user

type userDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}
