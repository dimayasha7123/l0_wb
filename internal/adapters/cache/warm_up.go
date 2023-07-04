package cache

import (
	"fmt"
	"strings"

	"l0_wb/internal/domain"
)

func (c cache) WarmUp(models []domain.Model) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	existed := make([]string, 0)

	for _, model := range models {
		_, ok := c.data[model.OrderUID]
		if ok {
			existed = append(existed, model.OrderUID)
			continue
		}

		c.data[model.OrderUID] = model
	}

	if len(existed) != 0 {
		return fmt.Errorf("found models with duplicate ids: %v", strings.Join(existed, ", "))
	}

	return nil
}
