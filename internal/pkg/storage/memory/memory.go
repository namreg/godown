package memory

import (
	"sync"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Storage represents a storage that store its all data in memory. Implements Storage interaface
type Storage struct {
	mu    sync.RWMutex
	items map[storage.Key]*storage.Value
}

//New creates a new memory storage
func New() *Storage {
	return &Storage{
		items: make(map[storage.Key]*storage.Value),
	}
}

//Put puts a new value that will be returned by ValueSetter
func (strg *Storage) Put(key storage.Key, setter storage.ValueSetter) error {
	strg.mu.Lock()

	value := strg.items[key]

	var err error

	if value, err = setter(value); err == nil {
		if value == nil {
			delete(strg.items, key)
		} else {
			strg.items[key] = value
		}
	}

	strg.mu.Unlock()

	return err
}

//Get gets a value of the storage by the given Key
func (strg *Storage) Get(key storage.Key) (*storage.Value, error) {
	strg.mu.RLock()
	defer strg.mu.RUnlock()

	if value, exists := strg.items[key]; exists {
		return value, nil
	}

	return nil, storage.ErrKeyNotExists
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

//All returns all stored keys
func (strg *Storage) All() (map[storage.Key]*storage.Value, error) {
	strg.mu.RLock()
	vals := strg.items
	strg.mu.RUnlock()
	return vals, nil
}
