package app

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
	"l0_wb/internal/utils/logger"
)

func (a Service) Get(ctx context.Context, uid string) (domain.Model, error) {
	model, err := a.cache.Get(ctx, uid)
	if err == nil {
		return model, nil
	}

	logger.Log().Infof("can't find model with uid = %s in cache: %v", uid, err)

	model, err = a.repository.Get(ctx, uid)
	if err != nil {
		return domain.Model{}, fmt.Errorf("can't find model with uid = %s in cache and repository: %v", uid, err)
	}

	err = a.cache.Add(ctx, model)
	if err != nil {
		logger.Log().Errorf("can't add model with uid = %s to cache: %v", uid, err)
	}

	return model, nil
}
