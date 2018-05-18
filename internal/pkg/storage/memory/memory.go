package memory

import (
	"sync"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Storage represents a storage that store its all data in memory. Implements Storage interaface
type Storage struct {
	mu    sync.RWMutex
	items map[storage.Key]storage.Value
}

//New creates a new memory storage
func New() *Storage {
	return &Storage{
		items: make(map[storage.Key]storage.Value),
	}
}

//Put puts a value to the storage by the given key
func (strg *Storage) Put(key storage.Key, value storage.Value) error {
	strg.mu.Lock()
	strg.items[key] = value
	strg.mu.Unlock()
	return nil
}

//Get gets a value of storage by the given key
func (strg *Storage) Get(key storage.Key) (storage.Value, error) {
	strg.mu.RLock()
	val := strg.items[key]
	strg.mu.RUnlock()
	return val, nil
}

//Del deletes a value by the given key
func (strg *Storage) Del(key storage.Key) error {
	strg.mu.Lock()
	delete(strg.items, key)
	strg.mu.Unlock()
	return nil
}

//Keys returns all stored keys
func (strg *Storage) Keys() ([]storage.Key, error) {
	strg.mu.RLock()

	keys := make([]storage.Key, 0, len(strg.items))
	for key := range strg.items {
		keys = append(keys, key)
	}

	strg.mu.RUnlock()
	return keys, nil
}
