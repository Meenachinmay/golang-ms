package memory

import (
	"context"
	"movieexample.com/metadata/internal/repository"
	"movieexample.com/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New create a new memory repository
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Get Retrieves movie metadata for by movie id
func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata
	return nil
}
