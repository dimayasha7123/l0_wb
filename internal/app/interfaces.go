package app

import (
	"context"

	"l0_wb/internal/domain"
)

type repository interface {
	GetAll(ctx context.Context) ([]domain.Model, error)
	Get(ctx context.Context, uid string) (domain.Model, error)
	Add(ctx context.Context, model domain.Model) error
	Exists(ctx context.Context, uid string) (bool, error)
}

type cache interface {
	WarmUp(models []domain.Model) error
	Get(ctx context.Context, uid string) (domain.Model, error)
	Add(ctx context.Context, model domain.Model) error
	Exists(ctx context.Context, uid string) (bool, error)
}

type validator interface {
	Validate(data []byte) error
}
