package postgresql

type StorageConfig struct {
	User       string `yaml:"user"`
	DBName     string `yaml:"dbname"`
	Password   string `env:"POSTGRES_PASSWORD"` // кому забудем передать пароль, не запустится
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	SSLmode    string `yaml:"sslmode"`
	SearchPath string `yaml:"search_path"`
}
