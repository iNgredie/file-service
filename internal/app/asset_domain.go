package app

import (
	"github.com/iNgredie/file-service/internal/asset/adapter/postgres"
	"github.com/iNgredie/file-service/internal/asset/controller/http_router"
	"github.com/iNgredie/file-service/internal/asset/usecase"
)

func AssetDomain(d Dependencies) {
	assetUseCase := usecase.New(
		postgres.New(d.Postgres.Pool),
	)
	http_router.AssetRouter(d.RouterHTTP, assetUseCase)
}
