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
	Mode       string
}

func New() (*Config, error) {
	cfg := &Config{}

	EnvToString(&cfg.HttpPort, "HTTP_PORT", ":8080")
	EnvToString(&cfg.DBHost, "DB_HOST", "ec2-52-212-228-71.eu-west-1.compute.amazonaws.com")
	EnvToString(&cfg.DBPort, "DB_PORT", "5432")
	EnvToString(&cfg.DBUsername, "DB_USERNAME", "fzcamrgntritxl")
	EnvToString(&cfg.DBName, "DB_NAME", "d7r6s29qsbs7ah")
	EnvToString(&cfg.DBPassword, "DB_PASSWORD", "test")
	EnvToString(&cfg.DBSSLMode, "DB_SSL_MODE", "require")

	EnvToString(&cfg.Mode, "MODE", "postgres")

	return cfg, nil
}
