package usecase

import (
	"context"
	"fmt"
	"github.com/iNgredie/file-service/internal/asset/dto"
)

func (u *UseCase) GetAsset(
	ctx context.Context,
	name string,
) (dto.GetAssetOutput, error) {

	asset, err := u.postgres.GetAsset(ctx, name)
	if err != nil {
		return dto.GetAssetOutput{}, fmt.Errorf("u.postgres.GetAsset: %w", err)
	}
	return dto.GetAssetOutput{
		Name:      asset.Name,
		Data:      asset.Data,
		CreatedAt: asset.CreatedAt,
	}, err
}
