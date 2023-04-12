package logging

type LoggingConfig struct {
	Dir        string `yaml:"dir"`
	Filename   string `yaml:"filename"`
	ProjectDir string `yaml:"project_dir"`
}
