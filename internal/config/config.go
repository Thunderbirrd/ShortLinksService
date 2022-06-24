package config

import (
	. "github.com/Thunderbirrd/ShortLinksService/pkg/utils"
)

type Config struct {
	HttpPort   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
	DBSSLMode  string
}

func New() (*Config, error) {
	cfg := &Config{}

	EnvToString(&cfg.HttpPort, "HTTP_PORT", ":8080")
	EnvToString(&cfg.DBHost, "DB_HOST", "127.0.0.1")
	EnvToString(&cfg.DBPort, "DB_PORT", "5432")
	EnvToString(&cfg.DBUsername, "DB_USERNAME", "user")
	EnvToString(&cfg.DBName, "DB_NAME", "test")
	EnvToString(&cfg.DBPassword, "DB_PASSWORD", "password")
	EnvToString(&cfg.DBSSLMode, "DB_SSL_MODE", "disable")

	return cfg, nil
}
