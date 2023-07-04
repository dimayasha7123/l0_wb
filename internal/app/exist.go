package app

import (
	"context"
	"fmt"

	"l0_wb/internal/utils/logger"
)

func (a Service) Exists(ctx context.Context, uid string) (bool, error) {
	exists, err := a.cache.Exists(ctx, uid)
	if err == nil {
		return exists, nil
	}

	logger.Log().Infof("can't find model with uid = %s in cache: %v", uid, err)

	exists, err = a.repository.Exists(ctx, uid)
	if err != nil {
		return false, fmt.Errorf("can't check existence of model with uid = %s in cache and repository: %v", uid, err)
	}

	if exists {
		model, err := a.repository.Get(ctx, uid)
		if err != nil {
			logger.Log().Errorf("can't find model with uid = %s in cache and repository but it was... : %v", uid, err)
		} else {
			err = a.cache.Add(ctx, model)
			if err != nil {
				logger.Log().Errorf("can't add model with uid = %s to cache: %v", uid, err)
			}
		}
	}

	return exists, nil
}
