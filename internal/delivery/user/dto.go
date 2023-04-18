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

// for swagger

type errorResponseDTO struct {
	Message string `json:"message"`
}

type tokenDTO struct {
	Csrf string `json:"csrf-token"`
}

type profileDTO struct {
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	AvatarURL string `json:"avatar_url"`
}
