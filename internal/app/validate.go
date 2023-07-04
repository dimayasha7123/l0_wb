package app

import (
	"context"
	"fmt"
)

func (a Service) Validate(ctx context.Context, data []byte) error {
	err := a.validator.Validate(data)
	if err != nil {
		return fmt.Errorf("data not valid: %v", err)
	}
	return nil
}
