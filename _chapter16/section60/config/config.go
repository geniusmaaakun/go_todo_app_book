package config

//環境変数から設定を取得

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}
	//構造体のフィールドを環境変数から取得し、設定
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
