package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string `validate:"required"`

	ServerPort         string `validate:"required"`
	ServerTimeoutRead  int    `validate:"required"`
	ServerTimeoutWrite int    `validate:"required"`
	ServerTimeoutIdle  int    `validate:"required"`

	DBHostname string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var config Config

	config.Env = os.Getenv("ENV")

	// Get server config
	config.ServerPort = os.Getenv("SERVER_PORT")
	timeout, err := strconv.Atoi(os.Getenv("SERVER_TIMEOUT_READ"))
	if err != nil {
		return nil, err
	}
	config.ServerTimeoutRead = timeout

	timeout, err = strconv.Atoi(os.Getenv("SERVER_TIMEOUT_WRITE"))
	if err != nil {
		return nil, err
	}
	config.ServerTimeoutWrite = timeout

	timeout, err = strconv.Atoi(os.Getenv("SERVER_TIMEOUT_IDLE"))
	if err != nil {
		return nil, err
	}
	config.ServerTimeoutIdle = timeout

	// Get database environment variables
	config.DBHostname = os.Getenv("DB_HOSTNAME")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBSSLMode = os.Getenv("DB_SSLMODE")

	return &config, nil
}
