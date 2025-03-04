package config

import "example.com/m/pkg/postgres"

type Config struct {
	DB postgres.DBConfig
}

func NewConfig() *Config {
	return nil
}
