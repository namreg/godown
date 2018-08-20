package memory

import (
	"sync"
	"time"

	"github.com/namreg/godown-v2/internal/clock"
	"github.com/namreg/godown-v2/internal/storage"
)

type memoryClock interface {
	Now() time.Time
}

//Storage represents a storage that store its all data in memory. Implements Storage interaface
type Storage struct {
	clck memoryClock

	metaMu sync.RWMutex
	meta   map[storage.MetaKey]storage.MetaValue

	mu           sync.RWMutex
	items        map[storage.Key]*storage.Value
	itemsWithTTL map[storage.Key]*storage.Value // items thats have ttl and will be processed by GC
}

//WithClock sets memoryClock implementation
func WithClock(clck memoryClock) func(*Storage) {
	return func(strg *Storage) {
		strg.clck = clck
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
		meta:         make(map[storage.MetaKey]storage.MetaValue),
		items:        items,
		itemsWithTTL: itemsWithTTL,
	}

	for _, f := range opts {
		f(strg)
	}

	if strg.clck == nil {
		strg.clck = clock.New()
	}

	return strg
}

//Put puts a new value that will be returned by ValueSetter.
func (strg *Storage) Put(key storage.Key, setter storage.ValueSetter) error {
	strg.mu.Lock()

	var value *storage.Value
	var exists bool
	var err error

	if value, exists = strg.items[key]; exists && value.IsExpired(strg.clck.Now()) {
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

	if value, exists := strg.items[key]; exists && !value.IsExpired(strg.clck.Now()) {
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
	defer strg.mu.RUnlock()

	keys := make([]storage.Key, 0, len(strg.items))
	for k, v := range strg.items {
		if !v.IsExpired(strg.clck.Now()) {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

//All returns all stored values
func (strg *Storage) All() (map[storage.Key]*storage.Value, error) {
	strg.mu.RLock()
	defer strg.mu.RUnlock()
	return strg.items, nil
}

//AllWithTTL returns all stored values thats have TTL
func (strg *Storage) AllWithTTL() (map[storage.Key]*storage.Value, error) {
	strg.mu.RLock()
	defer strg.mu.RUnlock()
	return strg.itemsWithTTL, nil
}

//PutMeta puts a meta value at the given key.
func (strg *Storage) PutMeta(key storage.MetaKey, value storage.MetaValue) error {
	strg.metaMu.Lock()
	strg.meta[key] = value
	strg.metaMu.Unlock()
	return nil
}

//GetMeta gets a meta value at the given key.
func (strg *Storage) GetMeta(key storage.MetaKey) (storage.MetaValue, error) {
	strg.metaMu.RLock()
	value := strg.meta[key]
	strg.metaMu.RUnlock()
	return value, nil
}
