package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	ttlJSONField   = "ttl"
	typeJSONField  = "type"
	valueJSONField = "value"
)

//UnmarshalJSON unmarshal a valid JSON to the value.
func (v *Value) UnmarshalJSON(j []byte) error {
	if len(j) == 0 {
		return nil
	}

	val := Value{}

	m := make(map[string]interface{})
	if err := json.Unmarshal(j, &m); err != nil {
		return err
	}

	t, ok := m[typeJSONField]
	if !ok {
		return fmt.Errorf("could not recognize type")
	}

	val.dataType = DataType(t.(string))

	for k, mv := range m {
		switch k {
		case ttlJSONField:
			f, ok := mv.(float64)
			if !ok {
				return fmt.Errorf("could not unmarshal ttl: ttl is not a float64")
			}
			val.ttl = int64(f)
		case valueJSONField:
			switch val.dataType {
			case StringDataType:
				s, ok := mv.(string)
				if !ok {
					return fmt.Errorf("could not unmarshal string value: not a string representation")
				}
				val.data = s
			case ListDataType:
				l, ok := mv.([]string)
				if !ok {
					return fmt.Errorf("could not unmarshal list value: not a list representation")
				}
				val.data = l
			case MapDataType:
				tm, ok := mv.(map[string]interface{})
				if !ok {
					return fmt.Errorf("could not unmarshal map value: not a map representation")
				}
				vmap := make(map[string]string, len(tm))
				for mkey, mvalue := range tm {
					sv, ok := mvalue.(string)
					if !ok {
						return fmt.Errorf("could not unmarshal map value: key %q is not a string", mkey)
					}
					vmap[mkey] = sv
				}
				val.data = vmap
			case BitMapDataType:
				l, ok := mv.([]uint64)
				if !ok {
					return fmt.Errorf("could not unmarshal bitmap value: not a bitmap representation")
				}
				val.data = l
			}
		}
	}

	*v = val

	return nil
}

//MarshalJSON marshal a value to the valid JSON.
func (v *Value) MarshalJSON() ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	buf := new(bytes.Buffer)
	buf.WriteByte('{')

	ttlBytes := make([]byte, 0)
	ttlBytes = strconv.AppendInt(ttlBytes, v.ttl, 10)
	buf.WriteString(`"ttl":`)
	buf.Write(ttlBytes)

	buf.WriteByte(',')

	buf.WriteString(`"type":"`)
	buf.WriteString(v.dataType.String())
	buf.WriteByte('"')

	buf.WriteByte(',')

	buf.WriteString(`"value":`)

	switch v.dataType {
	case StringDataType:
		buf.WriteByte('"')
		buf.WriteString(v.data.(string))
		buf.WriteByte('"')
	case ListDataType:
		buf.WriteByte('[')
		list := v.data.([]string)
		size := len(list)
		for i, lv := range list {
			buf.WriteByte('"')
			buf.WriteString(lv)
			buf.WriteByte('"')
			if i != size-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')
	case BitMapDataType:
		buf.WriteByte('[')
		list := v.data.([]uint64)
		size := len(list)
		for i, lv := range list {
			buf.Write(strconv.AppendUint([]byte{}, lv, 10))
			if i != size-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')
	case MapDataType:
		buf.WriteByte('{')
		m := v.data.(map[string]string)
		size := len(m)
		i := 0
		for key, val := range m {
			buf.WriteByte('"')
			buf.WriteString(key)
			buf.WriteString(`":`)
			buf.WriteByte('"')
			buf.WriteString(val)
			buf.WriteByte('"')
			i++
			if i != size {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
