package config

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	// для уровня логирования
	IsDebug bool `yaml:"is_debug" env-default:"true"`
	Server  struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"server"`
	Postgres postgresql.StorageConfig `yaml:"postgres"`
	Logging  logging.LoggingConfig    `yaml:"logging"`
}

var instance *Config

func GetCfg(configFile string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	instance = &Config{}
	err = cleanenv.ReadConfig(configFile, instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
