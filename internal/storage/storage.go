package storage

import (
	"errors"
	"time"
)

var (
	//ErrKeyNotExists means that key does not exist. Returns by GetKey method.
	ErrKeyNotExists = errors.New("storage: key does not exist")
)

//DataType represents type of data
type DataType string

const (
	//StringDataType is the string data type.
	StringDataType DataType = "string"
	//BitMapDataType is the bitmap data type. Stored as int64 integer.
	BitMapDataType DataType = "bitmap"
	//ListDataType is the list data type. Stored as slice of string.
	ListDataType DataType = "list"
	//MapDataType is the hash map type. Stored as map[string][string].
	MapDataType DataType = "map"
)

//DataType returns string representation of the DataType.
func (dt DataType) String() string {
	return string(dt)
}

//Key is the key of a value in the storage.
type Key string

//Value represents a single value of a storage.
type Value struct {
	ttl      int64 //unix time
	data     interface{}
	dataType DataType
}

//MetaKey is key of the meta value.
type MetaKey string

//MetaValue is a system system and not visible for users.
type MetaValue string

//ValueSetter guarantees transactional updates.
type ValueSetter func(old *Value) (new *Value, err error)

//Data returns data of the value.
func (v *Value) Data() interface{} {
	return v.data
}

//Type returns a DataType of the value.
func (v *Value) Type() DataType {
	return v.dataType
}

//IsExpired indicates whether the value is expired.
func (v *Value) IsExpired(till time.Time) bool {
	if v.ttl < 0 {
		return false
	}
	return v.ttl < till.Unix()
}

//SetTTL sets expiration time of the value.
func (v *Value) SetTTL(at time.Time) {
	v.ttl = at.Unix()
}

//TTL returns expired time of the value.
func (v *Value) TTL() int64 {
	return v.ttl
}

//NewString creates a new value with StringDataType.
func NewString(str string) *Value {
	return &Value{
		data:     str,
		dataType: StringDataType,
		ttl:      -1,
	}
}

//NewBitMap creates a new value of the BitMapDataType. Stored as uint64 integer.
func NewBitMap(value []uint64) *Value {
	return &Value{
		data:     value,
		dataType: BitMapDataType,
		ttl:      -1,
	}
}

//NewList creates a new value of the ListDataType. Stored as slice of strings.
func NewList(data []string) *Value {
	return &Value{
		data:     data,
		dataType: ListDataType,
		ttl:      -1,
	}
}

//NewMap creates a new value of the MapDataType.
func NewMap(val map[string]string) *Value {
	return &Value{
		data:     val,
		dataType: MapDataType,
		ttl:      -1,
	}
}
