package storage

import (
	"errors"
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
	expiredAt int64 //unix time
	data      interface{}
	dataType  DataType
}

//Data returns data of the value
func (v *Value) Data() interface{} {
	return v.data
}

//Type returns a DataType of the value
func (v *Value) Type() DataType {
	return v.dataType
}

//IsExpired indicates wheter the value is expired
func (v *Value) IsExpired() bool {
	return v.expiredAt > 0 && v.expiredAt < time.Now().Unix()
}

//NewStringValue creates a new value with StringDataType
func NewStringValue(str string) *Value {
	return &Value{
		data:     str,
		dataType: StringDataType,
	}
}

//NewBitMapValue creates a new value of the BitMapDataType. Stored as int64 integer
func NewBitMapValue(value int64) *Value {
	return &Value{
		data:     value,
		dataType: BitMapDataType,
	}
}

//NewListValue creates a new value of the ListDataType. Stored as slice of strings
func NewListValue(vals ...string) *Value {
	data := make([]string, 0, len(vals))
	data = append(data, vals...)
	return &Value{
		data:     data,
		dataType: ListDataType,
	}
}

//NewMapValue creates a new value of the MapDataType.
func NewMapValue(val map[string]string) *Value {
	return &Value{
		data:     val,
		dataType: MapDataType,
	}
}

//Storage represents a storage
type Storage interface {
	//Put puts a new value that will be returned by ValueSetter at the given Key
	Put(key Key, setter ValueSetter) error
	//Get gets a Value of a storage by the given Key.
	Get(key Key) (*Value, error)
	//Del deletes a value by the given Key
	Del(key Key) error
	//Keys returns all stored keys
	Keys() ([]Key, error)
	//All returns all stored values
	All() (map[Key]*Value, error)
}

//ValueSetter is a callback that calls by Storage during Puts a new Value
//Returns a new Value
type ValueSetter func(old *Value) (new *Value, err error)
