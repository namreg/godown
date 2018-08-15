package storage

import (
	"errors"
	"sync"
	"time"
)

var (
	//ErrKeyNotExists means that key does not exist. Returns by GetKey method
	ErrKeyNotExists = errors.New("storage: key does not exist")
)

//DataType represents type of data
type DataType string

const (
	//StringDataType is the string data type
	StringDataType DataType = "string"
	//BitMapDataType is the bitmap data type. Stored as int64 integer
	BitMapDataType DataType = "bitmap"
	//ListDataType is the list data type. Stored as slice of string
	ListDataType DataType = "list"
	//MapDataType is the hash map type. Stored as map[string][string]
	MapDataType DataType = "map"
)

//DataType returns string representation of the DataType
func (dt DataType) String() string {
	return string(dt)
}

//Key is the key of a value in the storage
type Key string

//Value represents a single value of a storage
type Value struct {
	ttl      int64 //unix time
	data     interface{}
	dataType DataType
}

//Data returns data of the value
func (v *Value) Data() interface{} {
	return v.data
}

//Type returns a DataType of the value
func (v *Value) Type() DataType {
	return v.dataType
}

//IsExpired indicates whether the value is expired
func (v *Value) IsExpired(till time.Time) bool {
	if v.ttl < 0 {
		return false
	}
	return v.ttl < till.Unix()
}

//SetTTL sets expiration time of the value
func (v *Value) SetTTL(at time.Time) {
	v.ttl = at.Unix()
}

//TTL returns expired time of the value
func (v *Value) TTL() int64 {
	return v.ttl
}

//NewStringValue creates a new value with StringDataType
func NewStringValue(str string) *Value {
	return &Value{
		data:     str,
		dataType: StringDataType,
		ttl:      -1,
	}
}

//NewBitMapValue creates a new value of the BitMapDataType. Stored as uint64 integer
func NewBitMapValue(value []uint64) *Value {
	return &Value{
		data:     value,
		dataType: BitMapDataType,
		ttl:      -1,
	}
}

//NewListValue creates a new value of the ListDataType. Stored as slice of strings
func NewListValue(data []string) *Value {
	return &Value{
		data:     data,
		dataType: ListDataType,
		ttl:      -1,
	}
}

//NewMapValue creates a new value of the MapDataType.
func NewMapValue(val map[string]string) *Value {
	return &Value{
		data:     val,
		dataType: MapDataType,
		ttl:      -1,
	}
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/pkg/storage.Storage -o ./ -s "_mock.go" -b test

//Storage represents a storage
type Storage interface {
	sync.Locker
	//RLock locks storage for reading.
	RLock()
	//RUnlock undoes a single RLock call.
	RUnlock()
	//Put puts a new *Value at the given Key.
	//Warning: method is not thread safe! You should call Lock mannually before calling
	Put(key Key, val *Value) error
	//Get gets a Value of a storage by the given Key.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
	Get(key Key) (*Value, error)
	//Del deletes a value by the given Key.
	//Warning: method is not thread safe! You should call Lock mannually before calling.
	Del(key Key) error
	//Keys returns all stored keys.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
	Keys() ([]Key, error)
	//All returns all stored values.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
	All() (map[Key]*Value, error)
	//AllWithTTL returns all stored values thats have TTL.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
	AllWithTTL() (map[Key]*Value, error)
}
