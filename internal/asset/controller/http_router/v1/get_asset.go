package v1

import (
	"errors"
	"github.com/iNgredie/file-service/internal/asset/entity"
	"github.com/iNgredie/file-service/pkg/render"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func (h *Handlers) GetAsset(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	var err error
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[len(parts)-1]
	output, err := h.uc.GetAsset(ctx, name)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			log.Error().Err(err).Msg("uc.GetAsset: not found")
			http.Error(w, "not found", http.StatusNotFound)

			return

		case errors.Is(err, entity.ErrUUIDInvalid), errors.Is(err, entity.ErrStatusInvalid):
			log.Error().Err(err).Msg("uc.GetAsset: validate error")
			http.Error(w, "validate error", http.StatusBadRequest)

			return

		default:
			log.Error().Err(err).Msg("uc.GetAsset: internal error")
			http.Error(w, "internal error", http.StatusInternalServerError)

			return
		}
	}

	render.JSON(w, output)
}
