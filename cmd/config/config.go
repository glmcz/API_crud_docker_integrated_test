package config

import (
	"simpleCloudService/internal/repository"
	"simpleCloudService/internal/serviceLayer"
)

type Config struct {
	ServerConfig   serviceLayer.ServerConfig `yaml:"server"`
	PostgresConfig repository.PostgresConfig `yaml:"postgres"`
	// tracer ...
}

func NewDefaultConfig() Config {
	var cfg Config

	return cfg
}
