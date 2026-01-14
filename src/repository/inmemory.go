package repository

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		store: make(map[string]string, 0),
	}
}

func (r *InMemoryRepository) GetURL(id string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	url := r.store[id]
	if url == "" {
		return "", errors.New("id not found in store")
	}

	return r.store[id], nil
}

func (r *InMemoryRepository) CreateURL(url string, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[id] = url

	return nil
}

func (r *InMemoryRepository) Exists(id string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	exists := r.store[id] != ""

	return exists, nil
}
