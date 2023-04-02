package setup

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	// для уровня логирования
	IsDebug bool `yaml:"is_debug" env-default:"true"`
	CORS    struct {
		AllowedOrigins string `yaml:"allowed_origins"`
	}
	Listen struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLmode  string `yaml:"sslmode"`
}

var instance *Config

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	instance = &Config{}
	err = cleanenv.ReadConfig("config.yml", instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
