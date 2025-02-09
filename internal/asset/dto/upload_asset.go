package dto

import (
	"github.com/iNgredie/file-service/internal/asset/entity"
)

type UploadAssetOutput struct {
	Name string `json:"name"`
}

type UploadAssetInput struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func (i *UploadAssetInput) Validate() error {
	if i.Name == "" {
		return entity.ErrNameInvalid
	}

	return nil
}
