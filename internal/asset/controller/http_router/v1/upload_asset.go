package v1

import (
	"encoding/json"
	"errors"
	"github.com/iNgredie/file-service/internal/asset/dto"
	"github.com/iNgredie/file-service/internal/asset/entity"
	"github.com/iNgredie/file-service/pkg/render"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) UploadAsset(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	input := dto.UploadAssetInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Error().Err(err).Msg("json.NewDecoder")
		http.Error(w, "json error", http.StatusBadRequest)

		return
	}

	err = input.Validate()
	if err != nil {
		log.Error().Err(err).Msg("input.Validate")
		http.Error(w, "validate error", http.StatusBadRequest)

		return
	}

	output, err := h.uc.UploadAsset(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			log.Error().Err(err).Msg("uc.UploadAsset: not found")
			http.Error(w, "not found", http.StatusNotFound)

			return

		case errors.Is(err, entity.ErrUUIDInvalid), errors.Is(err, entity.ErrStatusInvalid):
			log.Error().Err(err).Msg("uc.UploadAsset: validate error")
			http.Error(w, "validate error", http.StatusBadRequest)

			return

		default:
			log.Error().Err(err).Msg("uc.UploadAsset: internal error")
			http.Error(w, "internal error", http.StatusInternalServerError)

			return
		}
	}

	render.JSON(w, output)
}
