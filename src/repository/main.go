package repository

import (
	"sync"
)

type Repository interface {
	GetURL(id string) string
	CreateURL(url string, id string)
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		store: make(map[string]string, 0),
	}
}

func (r *InMemoryRepository) GetURL(id string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.store[id]
}

func (r *InMemoryRepository) CreateURL(url string, id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[id] = url
}
