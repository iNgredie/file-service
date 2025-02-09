package config

import (
	"fmt"
	"github.com/iNgredie/file-service/pkg/http_server"
	"github.com/iNgredie/file-service/pkg/logger"
	"github.com/iNgredie/file-service/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	HTTP     http_server.Config
	Postgres postgres.Config
	Logger   logger.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
