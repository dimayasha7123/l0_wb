package cache

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
)

func (c cache) Add(ctx context.Context, model domain.Model) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, ok := c.data[model.OrderUID]
	if ok {
		return fmt.Errorf("model with uid = %s already exists", model.OrderUID)
	}

	c.data[model.OrderUID] = model

	return nil
}
