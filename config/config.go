package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
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
	Storage postgresql.StorageConfig `yaml:"storage"`
	Redis   redis.RedisConfig        `yaml:"redis"`
}

var instance *Config

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	instance = &Config{}
	err = cleanenv.ReadConfig("config_prod.yml", instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
