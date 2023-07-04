package repository

import (
	"context"
	"fmt"
)

func (r repository) Exists(ctx context.Context, uid string) (bool, error) {
	query := `select exists(select 1 from models where order_uid = $1);`

	var exists bool
	err := r.pool.QueryRow(ctx, query, uid).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("can't check existence of model with id = %s: %v", uid, err)
	}

	return exists, nil
}
