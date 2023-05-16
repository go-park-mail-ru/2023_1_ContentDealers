package setup

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	contentGateway "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/session"
	userGateway "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user"
	favGateway "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user_action"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	// для уровня логирования
	ApiGateway struct {
		CORS struct {
			AllowedOrigins string `yaml:"allowed_origins"`
		}
		ServiceSession    session.ServiceSessionConfig        `yaml:"service_session"`
		ServiceUser       userGateway.ServiceUserConfig       `yaml:"service_user"`
		ServiceUserAction favGateway.ServiceUserActionConfig  `yaml:"service_user_action"`
		ServiceContent    contentGateway.ServiceContentConfig `yaml:"service_content"`
		Server            struct {
			BindIP            string `yaml:"bind_ip"`
			Port              string `yaml:"port" env-default:"8080"`
			WriteTimeout      int    `yaml:"write_timeout"`
			ReadTimeout       int    `yaml:"read_timeout"`
			ReadHeaderTimeout int    `yaml:"read_header_timeout"`
			ShutdownTimeout   int    `yaml:"shutdown_timeout"`
		} `yaml:"server"`
		Avatar  user.AvatarConfig
		CSRF    csrf.CSRFConfig       `yaml:"csrf"`
		Logging logging.LoggingConfig `yaml:"logging"`
	} `yaml:"api_gateway"`
	Content struct {
		Server struct {
			BindIP string `yaml:"bind_ip"`
			Port   string `yaml:"port" env-default:"8080"`
		} `yaml:"server"`
		Postgres postgresql.StorageConfig `yaml:"postgres"`
		Logging  logging.LoggingConfig    `yaml:"logging"`
	} `yaml:"content"`
	Favorites struct {
		Server struct {
			BindIP string `yaml:"bind_ip"`
			Port   string `yaml:"port" env-default:"8080"`
		} `yaml:"server"`
		Postgres postgresql.StorageConfig `yaml:"postgres"`
		Logging  logging.LoggingConfig    `yaml:"logging"`
		Views    struct {
			ThresholdViewProgress float32 `yaml:"threshold_view_progress"`
		} `yaml:"views"`
	} `yaml:"favorites"`
	Session struct {
		Server struct {
			BindIP string `yaml:"bind_ip"`
			Port   string `yaml:"port" env-default:"8080"`
		} `yaml:"server"`
		Session struct {
			ExpiresAt int `yaml:"expires_at"`
		} `yaml:"session"`
		Redis   redis.RedisConfig     `yaml:"redis"`
		Logging logging.LoggingConfig `yaml:"logging"`
	} `yaml:"session"`
	User struct {
		Server struct {
			BindIP string `yaml:"bind_ip"`
			Port   string `yaml:"port" env-default:"8080"`
		} `yaml:"server"`
		Postgres postgresql.StorageConfig `yaml:"postgres"`
		Logging  logging.LoggingConfig    `yaml:"logging"`
	} `yaml:"user"`
}

var instance *Config

func GetCfg(configFile string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	instance = &Config{}
	err = cleanenv.ReadConfig(configFile, instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
