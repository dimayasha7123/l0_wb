package repository

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
)

func (r repository) Get(ctx context.Context, uid string) (domain.Model, error) {
	query := `
	select *
	from models
	where order_uid = $1;
	`

	row := r.pool.QueryRow(ctx, query, uid)
	ret, err := scanUnpackModel(row)
	if err != nil {
		return domain.Model{}, fmt.Errorf("can't scan and unpack model with uid = %s: %v", uid, err)
	}

	return ret, nil
}
