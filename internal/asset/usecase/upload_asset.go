package usecase

import (
	"context"
	"fmt"
	"github.com/iNgredie/file-service/internal/asset/dto"
	"github.com/iNgredie/file-service/internal/asset/entity"
)

func (u *UseCase) UploadAsset(
	ctx context.Context,
	input dto.UploadAssetInput,
) (dto.UploadAssetOutput, error) {

	var output dto.UploadAssetOutput

	a := entity.Asset{
		Name: input.Name,
		Data: input.Data,
	}

	err := u.postgres.UploadAsset(ctx, a)
	if err != nil {
		return output, fmt.Errorf("u.postgres.UploadAsset: %w", err)
	}
	return output, nil
}
