package cache

import (
	"sync"

	"l0_wb/internal/domain"
)

const defaultCapacity = 10000

type cache struct {
	mutex *sync.RWMutex
	data  map[string]domain.Model
}

func New() cache {
	return cache{
		mutex: &sync.RWMutex{},
		data:  make(map[string]domain.Model, defaultCapacity),
	}
}
