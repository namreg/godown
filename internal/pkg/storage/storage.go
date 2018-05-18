package storage

//DataType represents type of data
type DataType string

const (
	//StringDataType is the string data type
	StringDataType DataType = "string"
)

//DataType returns string representation of the DataType
func (dt DataType) String() string {
	return string(dt)
}

//IsEqual compares DataType with other
func (dt DataType) IsEqual(other DataType) bool {
	return string(dt) == string(other)
}

//Key is the key of value in the storage
type Key struct {
	value    string
	dataType DataType
}

//NewStringKey creates a new key with StringDataType
func NewStringKey(key string) Key {
	return Key{
		value:    key,
		dataType: StringDataType,
	}
}

//Value represents a value of storage
type Value interface{}

//Storage represents a storage
type Storage interface {
	//Put puts a value to the storage by the given key
	Put(Key, Value) error
	//Get gets a value of storage by the given key
	Get(Key) (Value, error)
	//Del deletes a value by the given key
	Del(Key) error
	//Keys returns all stored keys
	Keys() ([]Key, error)
}
