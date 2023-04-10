package redis

type RedisConfig struct {
	User  string `yaml:"user"`
	DBNum string `yaml:"dbnum"`
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
}
