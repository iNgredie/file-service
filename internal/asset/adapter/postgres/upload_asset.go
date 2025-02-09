package postgres

import (
	"context"
	"fmt"
	"github.com/iNgredie/file-service/internal/asset/entity"
)

func (p *Postgres) UploadAsset(
	ctx context.Context,
	a entity.Asset,
) (err error) {
	query := `
        INSERT INTO asset (id, name, data)
        VALUES ($1, $2, $3)
    `
	params := []any{
		a.Name,
		a.Data,
	}

	result, err := p.pool.Exec(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}
