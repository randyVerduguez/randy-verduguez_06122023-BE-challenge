package configs

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort int    `envconfig:"SERVER_PORT"`
	ServerIP   string `envconfig:"SERVER_IP"`
	ApiKey     string `envconfig:"API_KEY"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

func NewParsedConfig() (Config, error) {
	env_err := godotenv.Load(".env")

	if env_err != nil {
		env_err = fmt.Errorf("Error loading .env file %w", env_err)
		return Config{}, env_err
	}

	config := Config{}
	err := envconfig.Process("", &config)

	return config, err
}
