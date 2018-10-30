package storage

import (
	"reflect"
	"testing"
)

func TestMarshaler_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name   string
		input  *Value
		output []byte
		err    error
	}{
		{
			"nil value gets nil []byte",
			nil,
			nil,
			nil,
		},
		{
			"empty value gets JSON object with empty properties",
			&Value{},
			[]byte("{\"ttl\":0,\"type\":\"\",\"value\":}"),
			nil,
		},
		{
			"value with ttl gets JSON object with not empty ttl",
			&Value{ttl: 5},
			[]byte("{\"ttl\":5,\"type\":\"\",\"value\":}"),
			nil,
		},
		{
			"data set according to dataType of type string",
			&Value{dataType: StringDataType, data: "data"},
			[]byte("{\"ttl\":0,\"type\":\"string\",\"value\":\"data\"}"),
			nil,
		},
		{
			"data set to empty string array according to dataType of type list",
			&Value{dataType: ListDataType, data: []string{}},
			[]byte("{\"ttl\":0,\"type\":\"list\",\"value\":[]}"),
			nil,
		},
		{
			"data set to string array of 2 elements according to dataType of type list",
			&Value{dataType: ListDataType, data: []string{"str1", "str2"}},
			[]byte("{\"ttl\":0,\"type\":\"list\",\"value\":[\"str1\",\"str2\"]}"),
			nil,
		},
		{
			"data set to empty uint64 array according to dataType of type bitmap",
			&Value{dataType: BitMapDataType, data: []uint64{}},
			[]byte("{\"ttl\":0,\"type\":\"bitmap\",\"value\":[]}"),
			nil,
		},
		{
			"data set to uint64 array of 4 elements according to dataType of type bitmap",
			&Value{dataType: BitMapDataType, data: []uint64{1, 2, 3, 4}},
			[]byte("{\"ttl\":0,\"type\":\"bitmap\",\"value\":[1,2,3,4]}"),
			nil,
		},
		{
			"data set to empty map of string:string according to dataType of type map",
			&Value{dataType: MapDataType, data: map[string]string{}},
			[]byte("{\"ttl\":0,\"type\":\"map\",\"value\":{}}"),
			nil,
		},
		{
			"data set to a map of string:string of 1 element according to dataType of type map",
			&Value{dataType: MapDataType, data: map[string]string{"key1": "val1"}},
			[]byte("{\"ttl\":0,\"type\":\"map\",\"value\":{\"key1\":\"val1\"}}"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := tc.input.MarshalJSON()
			if tc.err != err {
				t.Errorf("given error %v doesn't equal to expected error %v", err, tc.err)
			}
			if !reflect.DeepEqual(m, tc.output) {
				t.Errorf("given value %v doesn't equal to expected value %v", string(m), string(tc.output))
			}
		})
	}
}

func TestValue_MarshalJSON_ShouldPanic(t *testing.T) {
	testCases := []struct {
		name  string
		input *Value
	}{
		{
			"dataType set but data nil should panic because of type conversion to string",
			&Value{dataType: StringDataType},
		},
		{
			"data of type int is not matching the dataType string => should panic because of type conversion to string",
			&Value{dataType: StringDataType, data: 0},
		},
		{
			"dataType is list but data contains []int => should panic because of type conversion to []string",
			&Value{dataType: ListDataType, data: []int{1, 2}},
		},
		{
			"dataType is bitmap but data contains []int => should panic because of type conversion to []uint64",
			&Value{dataType: BitMapDataType, data: []int{1, 2}},
		},
		{
			"dataType is map but data contains []int => should panic because of type conversion to map[string]string",
			&Value{dataType: BitMapDataType, data: []int{1, 2}},
		},
		{
			"dataType is map but data contains map[int]int => should panic because of type conversion to map[string]string",
			&Value{dataType: BitMapDataType, data: map[int]int{1: 2}},
		},
		{
			"dataType is map but data contains map[string]int => should panic because of type conversion to map[string]string",
			&Value{dataType: BitMapDataType, data: map[string]int{"1": 2}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assertPanic(t, func() { tc.input.MarshalJSON() })
		})
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func BenchmarkMarshaler_MarshalJSON(b *testing.B) {
	benchCases := []struct {
		name  string
		value *Value
	}{
		{
			name:  "nil value",
			value: nil,
		},
		{
			name:  "empty value",
			value: &Value{},
		},
		{
			name:  "string value",
			value: &Value{dataType: StringDataType, data: "a string to be marshalled"},
		},
		{
			name:  "list value",
			value: &Value{dataType: ListDataType, data: []string{"first string to be marshalled", "second string to be marshalled"}},
		},
		{
			name:  "bitmap value",
			value: &Value{dataType: BitMapDataType, data: []uint64{1, 2, 3, 4, 5}},
		},
		{
			name:  "map value",
			value: &Value{dataType: MapDataType, data: map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"}},
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bc.value.MarshalJSON()
			}
		})
	}
}
