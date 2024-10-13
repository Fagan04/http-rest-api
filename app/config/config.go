package config

type Config struct {
	BindAddress string `toml:"BIND_ADDRESS"`
	LogLevel    string `toml:"LOG_LEVEL"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
