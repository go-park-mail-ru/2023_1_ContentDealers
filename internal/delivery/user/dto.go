package user

type userCreateDTO struct {
	Id        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
	Birthday  string `json:"birthday"`
}

type userCredentialsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
