package postgresql

type StorageConfig struct {
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLmode  string `yaml:"sslmode"`
}
