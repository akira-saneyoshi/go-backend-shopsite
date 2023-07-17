package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Env        string `env:"GO_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"8000"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}