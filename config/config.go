package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Env        string `env:"GO_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"8000"`
	DBHost     string `env:"GO_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"GO_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"GO_DB_USER" envDefault:"gof1ber"`
	DBPassword string `env:"GO_DB_PASSWORD" envDefault:"gof1ber"`
	DBName     string `env:"GO_DB_NAME" envDefault:"shop_site"`
	RedisHost  string `env:"GO_REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort  int    `env:"GO_REDIS_PORT" envDefault:"36379"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}