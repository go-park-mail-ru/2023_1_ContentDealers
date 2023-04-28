package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	favGateway "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/favorites"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/session"
	userGateway "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	// для уровня логирования
	IsDebug bool `yaml:"is_debug" env-default:"true"`
	CORS    struct {
		AllowedOrigins string `yaml:"allowed_origins"`
	}
	ServiceSession   session.ServiceSessionConfig      `yaml:"service_sesion"`
	ServiceUser      userGateway.ServiceUserConfig     `yaml:"service_user"`
	ServiceFavorites favGateway.ServiceFavoritesConfig `yaml:"service_favorites"`
	Server           struct {
		BindIP            string `yaml:"bind_ip"`
		Port              string `yaml:"port" env-default:"8080"`
		WriteTimeout      int    `yaml:"write_timeout"`
		ReadTimeout       int    `yaml:"read_timeout"`
		ReadHeaderTimeout int    `yaml:"read_header_timeout"`
		ShutdownTimeout   int    `yaml:"shutdown_timeout"`
	} `yaml:"server"`
	Avatar      user.AvatarConfig
	CSRF        csrf.CSRFConfig          `yaml:"csrf"`
	Postgres    postgresql.StorageConfig `yaml:"postgres"`
	Redis       redis.RedisConfig        `yaml:"redis"`
	Logging     logging.LoggingConfig    `yaml:"logging"`
	ContentAddr string                   `content_addr`
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
