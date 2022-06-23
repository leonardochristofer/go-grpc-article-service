package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName    string `env:"SERVICE_NAME"`
	ServiceVersion string `env:"SERVICE_VERSION"`
	Env            string `env:"ENV"`
	ServicePort    string `env:"PORT"`
}

func Setup() (*Config, error) {
	cfgFile := "config.toml"

	if err := godotenv.Load(cfgFile); err != nil {
		return nil, err
	}

	cfg := &Config{}

	opts := &env.Options{TagName: "env"}

	if err := env.Parse(cfg, *opts); err != nil {
		return nil, err
	}

	return cfg, nil
}
