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
	//EnvToString(&cfg.DBHost, "DB_HOST", "127.0.0.1")
	//EnvToString(&cfg.DBPort, "DB_PORT", "5432")
	//EnvToString(&cfg.DBUsername, "DB_USERNAME", "user")
	//EnvToString(&cfg.DBName, "DB_NAME", "test")
	//EnvToString(&cfg.DBPassword, "DB_PASSWORD", "password")
	//EnvToString(&cfg.DBSSLMode, "DB_SSL_MODE", "disable")

	EnvToString(&cfg.DBHost, "DB_HOST", "ec2-52-212-228-71.eu-west-1.compute.amazonaws.com")
	EnvToString(&cfg.DBPort, "DB_PORT", "5432")
	EnvToString(&cfg.DBUsername, "DB_USERNAME", "fzcamrgntritxl")
	EnvToString(&cfg.DBName, "DB_NAME", "d7r6s29qsbs7ah")
	EnvToString(&cfg.DBPassword, "DB_PASSWORD", "7945968eed261e91271a72e87735b1f85715d098b45847238aaeae55b2d9d535")
	EnvToString(&cfg.DBSSLMode, "DB_SSL_MODE", "require")

	return cfg, nil
}
