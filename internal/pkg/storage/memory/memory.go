package memory

import (
	"sync"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Storage represents a storage that store its all data in memory. Implements Storage interaface
type Storage struct {
	mu           sync.RWMutex
	items        map[storage.Key]*storage.Value
	itemsWithTTL map[storage.Key]*storage.Value // items thats have ttl and will be processed by GC
}

//New creates a new memory storage with the given items
func New(items map[storage.Key]*storage.Value) *Storage {
	if items == nil {
		items = make(map[storage.Key]*storage.Value)
	}

	itemsWithTTL := make(map[storage.Key]*storage.Value)
	for k, v := range items {
		if v.TTL() > 0 {
			itemsWithTTL[k] = v
		}
	}
	return &Storage{
		items:        items,
		itemsWithTTL: itemsWithTTL,
	}
}

//Put puts a new value that will be returned by ValueSetter
func (strg *Storage) Put(key storage.Key, setter storage.ValueSetter) error {
	strg.mu.Lock()

	var value *storage.Value
	var exists bool
	var err error

	if value, exists = strg.items[key]; exists && value.IsExpired() {
		value = nil
	}

	if value, err = setter(value); err == nil {
		if value == nil {
			delete(strg.items, key)
			delete(strg.itemsWithTTL, key)
		} else {
			strg.items[key] = value
			if value.TTL() != -1 {
				strg.itemsWithTTL[key] = value
			}
		}
	}

	strg.mu.Unlock()

	return err
}

//Get gets a value of the storage by the given Key
func (strg *Storage) Get(key storage.Key) (*storage.Value, error) {
	strg.mu.RLock()
	defer strg.mu.RUnlock()

	if value, exists := strg.items[key]; exists && !value.IsExpired() {
		return value, nil
	}

	return nil, storage.ErrKeyNotExists
}

//Del deletes a value by the given key
func (strg *Storage) Del(key storage.Key) error {
	strg.mu.Lock()
	delete(strg.items, key)
	delete(strg.itemsWithTTL, key)
	strg.mu.Unlock()
	return nil
}

//Keys returns all stored keys
func (strg *Storage) Keys() ([]storage.Key, error) {
	strg.mu.RLock()

	keys := make([]storage.Key, 0, len(strg.items))
	for k, v := range strg.items {
		if !v.IsExpired() {
			keys = append(keys, k)
		}
	}

	strg.mu.RUnlock()
	return keys, nil
}

//All returns all stored values
func (strg *Storage) All() (map[storage.Key]*storage.Value, error) {
	strg.mu.RLock()
	vals := strg.items
	strg.mu.RUnlock()
	return vals, nil
}

//AllWithTTL returns all stored values thats have TTL
func (strg *Storage) AllWithTTL() (map[storage.Key]*storage.Value, error) {
	strg.mu.RLock()
	vals := strg.itemsWithTTL
	strg.mu.RUnlock()
	return vals, nil
}
