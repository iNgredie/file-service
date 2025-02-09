package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/iNgredie/file-service/internal/asset/entity"
)

func (p *Postgres) GetAsset(
	ctx context.Context,
	name string,
) (entity.Asset, error) {
	query := `
        SELECT name, data, created_at
        FROM asset
        WHERE name = $1
    `
	var asset entity.Asset
	err := p.pool.QueryRow(ctx, query, name).Scan(
		&asset.Name,
		&asset.Data,
		&asset.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Asset{}, entity.ErrNotFound
		}
		return entity.Asset{}, fmt.Errorf("failed to get asset: %w", err)
	}

	return entity.Asset{
		Name:      name,
		Data:      asset.Data,
		CreatedAt: asset.CreatedAt,
	}, nil
}
