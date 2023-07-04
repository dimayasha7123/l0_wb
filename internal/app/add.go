package app

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
)

func (a Service) Add(ctx context.Context, model domain.Model) error {
	err := a.repository.Add(ctx, model)
	if err != nil {
		return fmt.Errorf("can't add model with uid = %s to repository: %v", model.OrderUID, err)
	}

	err = a.cache.Add(ctx, model)
	if err != nil {
		return fmt.Errorf("can't add model with uid = %s to cache: %v", model.OrderUID, err)
	}

	return nil
}
