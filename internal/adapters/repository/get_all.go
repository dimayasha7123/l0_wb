package repository

import (
	"context"
	"errors"
	"fmt"

	"l0_wb/internal/domain"
)

func (r repository) GetAll(ctx context.Context) ([]domain.Model, error) {
	query := `select * from models`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't select models: %v", err)
	}
	defer rows.Close()

	ret := make([]domain.Model, 0)
	errs := make([]error, 0)
	for rows.Next() {
		model, err := scanUnpackModel(rows)
		if err != nil {
			errs = append(errs, fmt.Errorf("can't scan and unpack model: %v", err))
			continue
		}
		ret = append(ret, model)
	}

	if len(errs) != 0 {
		return nil, fmt.Errorf("can't scan all models: %v", errors.Join(errs...))
	}

	return ret, nil
}
