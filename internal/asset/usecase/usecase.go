package usecase

import (
	"context"
	"github.com/iNgredie/file-service/internal/asset/entity"
)

type Postgres interface {
	UploadAsset(
		ctx context.Context,
		a entity.Asset,
	) (err error)
	GetAsset(
		ctx context.Context,
		name string,
	) (entity.Asset, error)
}

type UseCase struct {
	postgres Postgres
}

func New(
	p Postgres,
) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
