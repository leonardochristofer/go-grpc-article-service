package config

import (
	"gopkg.in/gcfg.v1"
)

func Setup() (*Config, error) {
	cfgFile := "config.toml"

	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, cfgFile)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
