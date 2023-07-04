package cache

import (
	"context"
)

func (c cache) Exists(ctx context.Context, uid string) (bool, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	_, ok := c.data[uid]
	return ok, nil
}
