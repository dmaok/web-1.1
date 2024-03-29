package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
	SessionsKey string `toml:"sessions_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		DatabaseUrl: "host=localhost dbname=web_1_develop sslmode=disable",
	}
}
