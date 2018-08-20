package storage

import (
	"reflect"
	"testing"
	"time"

	"github.com/namreg/godown-v2/internal/clock"
	"github.com/stretchr/testify/assert"
)

func TestDataType_String(t *testing.T) {
	tests := []struct {
		name string
		dt   DataType
		want string
	}{
		{"StringDataType", StringDataType, "string"},
		{"BitMapDataType", BitMapDataType, "bitmap"},
		{"ListDataType", ListDataType, "list"},
		{"MapDataType", MapDataType, "map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.String(); got != tt.want {
				t.Errorf("DataType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Data_And_Value_Type(t *testing.T) {
	type fields struct {
		data     interface{}
		dataType DataType
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"string", fields{data: "string", dataType: StringDataType}, "string"},
		{"bitmap", fields{data: 1024, dataType: BitMapDataType}, 1024},
		{"list", fields{data: []string{"1", "2"}, dataType: ListDataType}, []string{"1", "2"}},
		{"map", fields{data: map[string]string{"key": "value"}, dataType: MapDataType}, map[string]string{"key": "value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value{
				data:     tt.fields.data,
				dataType: tt.fields.dataType,
			}
			if got := v.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value.Data() = %v, want %v", got, tt.want)
			}
			if got := v.Type(); !reflect.DeepEqual(got, tt.fields.dataType) {
				t.Errorf("Value.Type() = %v, want %v", got, tt.fields.dataType)
			}
		})
	}
}

func TestValue_IsExpired(t *testing.T) {
	tests := []struct {
		name  string
		value *Value
		want  bool
	}{
		{"no_ttl", &Value{ttl: -1}, false},
		{"not_expired", &Value{ttl: time.Now().Add(1 * time.Second).Unix()}, false},
		{"expired", &Value{ttl: time.Now().Add(-1 * time.Second).Unix()}, true},
	}
	clck := clock.New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.IsExpired(clck.Now()); got != tt.want {
				t.Errorf("Value.IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_SetTTL(t *testing.T) {
	value := &Value{
		ttl:      -1,
		data:     "string",
		dataType: StringDataType,
	}
	ttl := time.Now()
	value.SetTTL(ttl)

	assert.Equal(t, ttl.Unix(), value.ttl, "value.ttl is not set")
}

func TestValue_TTL(t *testing.T) {
	ttl := time.Now().Unix()
	value := &Value{
		ttl:      ttl,
		data:     "string",
		dataType: StringDataType,
	}

	assert.Equal(t, ttl, value.TTL())
}

func TestNewStringValue(t *testing.T) {
	value := NewStringValue("test string")
	assert.Equal(t, "test string", value.data)
	assert.Equal(t, StringDataType, value.dataType)
	assert.Equal(t, int64(-1), value.ttl)
}

func TestNewBitMapValue(t *testing.T) {
	value := NewBitMapValue([]uint64{35})
	assert.Equal(t, uint64(35), value.data.([]uint64)[0])
	assert.Equal(t, BitMapDataType, value.dataType)
	assert.Equal(t, int64(-1), value.ttl)
}

func TestNewListValue(t *testing.T) {
	value := NewListValue([]string{"test 1", "test 2", "test 3"})
	assert.Equal(t, []string{"test 1", "test 2", "test 3"}, value.data)
	assert.Equal(t, ListDataType, value.dataType)
	assert.Equal(t, int64(-1), value.ttl)
}

func TestNewMapValue(t *testing.T) {
	value := NewMapValue(map[string]string{"key": "value"})
	assert.Equal(t, map[string]string{"key": "value"}, value.data)
	assert.Equal(t, MapDataType, value.dataType)
	assert.Equal(t, int64(-1), value.ttl)
}
