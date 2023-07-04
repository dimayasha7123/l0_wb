package app

import (
	"context"
	"fmt"
)

type Service struct {
	cache      cache
	repository repository
	validator  validator
}

func New(ctx context.Context, cache cache, repository repository, validator validator) (Service, error) {
	models, err := repository.GetAll(ctx)
	if err != nil {
		return Service{}, fmt.Errorf("can't get all models from repository: %v", err)
	}

	err = cache.WarmUp(models)
	if err != nil {
		return Service{}, fmt.Errorf("can't warm up cache: %v", err)
	}

	return Service{
		cache:      cache,
		repository: repository,
		validator:  validator,
	}, nil
}
