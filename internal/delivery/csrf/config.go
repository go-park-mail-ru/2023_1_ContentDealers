package csrf

type CSRFConfig struct {
	Header    string `yaml:"header"`
	ExpiresAt int    `yaml:"expires_at"`
}
