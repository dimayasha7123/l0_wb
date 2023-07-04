package cache

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
)

func (c cache) Get(ctx context.Context, uid string) (domain.Model, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ret, ok := c.data[uid]
	if !ok {
		return domain.Model{}, fmt.Errorf("no model with uid = %s", uid)
	}

	return ret, nil
}
