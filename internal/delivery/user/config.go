package user

type AvatarConfig struct {
	NameFormFile string `yaml:"name_form_file"`
	MaxSizeBody  int    `yaml:"max_size_body"`
}
