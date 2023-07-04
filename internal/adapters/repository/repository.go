package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (repository, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return repository{}, fmt.Errorf("can't create pgxpool: %v", err)
	}

	return repository{pool: pool}, nil
}
