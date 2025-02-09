package app

import (
	"context"
	"fmt"
	"github.com/iNgredie/file-service/config"
	"github.com/iNgredie/file-service/pkg/http_server"
	"github.com/iNgredie/file-service/pkg/postgres"
	"github.com/iNgredie/file-service/pkg/router"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Dependencies struct {
	// Adapters
	Postgres *postgres.Pool

	// Controllers
	RouterHTTP *http.ServeMux
}

func Run(
	ctx context.Context,
	c config.Config,
) (err error) {
	var deps Dependencies

	// Adapters
	deps.Postgres, err = postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()
	// Controllers
	deps.RouterHTTP = router.New()
	// Domains
	AssetDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	log.Info().Msg("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		log.Info().Msg("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		log.Error().Err(err).Msg("App got notify: httpServer.Notify")
	}

	log.Info().Msg("App is stopping...")
}
