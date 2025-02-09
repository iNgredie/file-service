package http_router

import (
	ver1 "github.com/iNgredie/file-service/internal/asset/controller/http_router/v1"
	"github.com/iNgredie/file-service/internal/asset/usecase"
	"net/http"
)

func AssetRouter(
	r *http.ServeMux,
	uc *usecase.UseCase,
) {
	v1 := ver1.New(uc)
	v1Router := http.NewServeMux()
	v1Router.HandleFunc("/upload-asset/{asset_name}", v1.UploadAsset)
	v1Router.HandleFunc("/asset/{asset_name}", v1.GetAsset)

	// Регистрируем v1Router с префиксом /api
	r.Handle("/api/", http.StripPrefix("/api", v1Router))

	//r.Handle("/api/v1/", v1Router) // this is good practice
	//r.Handle("/api/", v1Router)
}
