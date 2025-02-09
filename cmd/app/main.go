package main

import (
	"context"
	"github.com/iNgredie/file-service/config"
	"github.com/iNgredie/file-service/internal/app"
	"github.com/iNgredie/file-service/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	err = app.Run(ctx, c)
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}

	log.Info().Msg("App stopped!")
}
