package memory

import (
	"sync"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/pkg/clock"
)

//Storage represents a storage that store its all data in memory. Implements Storage interaface
type Storage struct {
	clock        clock.Clock
	mu           sync.RWMutex
	items        map[storage.Key]*storage.Value
	itemsWithTTL map[storage.Key]*storage.Value // items thats have ttl and will be processed by GC
}

//WithClock sets Clock implementation
func WithClock(clck clock.Clock) func(*Storage) {
	return func(strg *Storage) {
		strg.clock = clck
	}
}

//New creates a new memory storage with the given items
func New(items map[storage.Key]*storage.Value, opts ...func(*Storage)) *Storage {
	if items == nil {
		items = make(map[storage.Key]*storage.Value)
	}

	itemsWithTTL := make(map[storage.Key]*storage.Value)
	for k, v := range items {
		if v.TTL() > 0 {
			itemsWithTTL[k] = v
		}
	}
	strg := &Storage{
		items:        items,
		itemsWithTTL: itemsWithTTL,
	}

	for _, f := range opts {
		f(strg)
	}

	if strg.clock == nil {
		strg.clock = clock.TimeClock{}
	}

	return strg
}

//Lock locks storage for writing
func (strg *Storage) Lock() {
	strg.mu.Lock()
}

//Unlock undoes a single Lock call
func (strg *Storage) Unlock() {
	strg.mu.Unlock()
}

//RLock locks storage for reading
func (strg *Storage) RLock() {
	strg.mu.RLock()
}

//RUnlock undoes a single RLock call
func (strg *Storage) RUnlock() {
	strg.mu.RUnlock()
}

//Put puts a new *Value at the given Key
func (strg *Storage) Put(key storage.Key, val *storage.Value) error {
	strg.items[key] = val
	if val.TTL() != -1 {
		strg.itemsWithTTL[key] = val
	}
	return nil
}

//Get gets a value of the storage by the given Key
func (strg *Storage) Get(key storage.Key) (*storage.Value, error) {
	if value, exists := strg.items[key]; exists && !value.IsExpired(strg.clock.Now()) {
		return value, nil
	}
	return nil, storage.ErrKeyNotExists
}

//Del deletes a value by the given key
func (strg *Storage) Del(key storage.Key) error {
	delete(strg.items, key)
	delete(strg.itemsWithTTL, key)
	return nil
}

//Keys returns all stored keys
func (strg *Storage) Keys() ([]storage.Key, error) {
	keys := make([]storage.Key, 0, len(strg.items))
	for k, v := range strg.items {
		if !v.IsExpired(strg.clock.Now()) {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

//All returns all stored values
func (strg *Storage) All() (map[storage.Key]*storage.Value, error) {
	return strg.items, nil
}

//AllWithTTL returns all stored values thats have TTL
func (strg *Storage) AllWithTTL() (map[storage.Key]*storage.Value, error) {
	return strg.itemsWithTTL, nil
}
