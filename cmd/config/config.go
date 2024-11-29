package config

import (
	"simpleCloudService/internal/api"
	"simpleCloudService/internal/repository"
)

type Config struct {
	ServerConfig   api.ServerConfig          `yaml:"server"`
	PostgresConfig repository.PostgresConfig `yaml:"postgres"`
	// tracer ...
}

func NewDefaultConfig(service string) Config {
	var cfg Config

	return cfg
}
