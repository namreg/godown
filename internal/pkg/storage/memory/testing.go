// +build test

package memory

import "github.com/namreg/godown-v2/internal/pkg/storage"

type testStorage struct {
	*Storage
}

//NewTestStorage creates a new storage with the given items. Uses only for testing purpose
func NewTestStorage(items, itemsWithTTL map[storage.Key]*storage.Value) *testStorage {
	mem := New()
	mem.items = items
	mem.itemsWithTTL = itemsWithTTL
	return &testStorage{mem}
}

//Items return all stored items
func (strg *testStorage) Items() map[storage.Key]*storage.Value {
	return strg.items
}

//Items return all stored items thats have ttl
func (strg *testStorage) ItemsWithTTL() map[storage.Key]*storage.Value {
	return strg.itemsWithTTL
}
