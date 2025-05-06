package hio

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func OrderedSerialize(v interface{}) (string, error) {
	s := &OrderedSerializer{}
	if err := s.doEncode(reflect.ValueOf(v)); err != nil {
		return "", err
	}
	return s.builder.String(), nil
}

type OrderedSerializer struct {
	builder strings.Builder
}

func (s *OrderedSerializer) doEncode(v reflect.Value) error {
	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		s.builder.WriteString("null")
		return nil
	}

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			s.builder.WriteString("true")
		} else {
			s.builder.WriteString("false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s.builder.WriteString(fmt.Sprintf("%d", v.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.builder.WriteString(fmt.Sprintf("%d", v.Uint()))
	case reflect.Float32, reflect.Float64:
		s.builder.WriteString(fmt.Sprintf("%g", v.Float()))
	case reflect.String:
		s.doEncodeString(v.String())
	case reflect.Slice, reflect.Array:
		return s.doEncodeArray(v)
	case reflect.Map:
		return s.doEncodeMap(v)
	case reflect.Interface:
		return s.doEncode(v.Elem())
	case reflect.Struct:
		return s.doEncodeStruct(v)
	default:
		return fmt.Errorf("unsupported type: %v", v.Kind())
	}
	return nil
}

func (s *OrderedSerializer) doEncodeString(str string) {
	s.builder.WriteByte('"')
	for _, c := range str {
		switch c {
		case '"':
			s.builder.WriteString("\\\"")
		case '\\':
			s.builder.WriteString("\\\\")
		case '\b':
			s.builder.WriteString("\\b")
		case '\f':
			s.builder.WriteString("\\f")
		case '\n':
			s.builder.WriteString("\\n")
		case '\r':
			s.builder.WriteString("\\r")
		case '\t':
			s.builder.WriteString("\\t")
		default:
			if c < 0x20 {
				s.builder.WriteString(fmt.Sprintf("\\u%04x", c))
			} else {
				s.builder.WriteRune(c)
			}
		}
	}
	s.builder.WriteByte('"')
}

func (s *OrderedSerializer) doEncodeArray(v reflect.Value) error {
	s.builder.WriteByte('[')
	length := v.Len()
	for i := 0; i < length; i++ {
		if i > 0 {
			s.builder.WriteByte(',')
		}
		if err := s.doEncode(v.Index(i)); err != nil {
			return err
		}
	}
	s.builder.WriteByte(']')
	return nil
}

func (s *OrderedSerializer) doEncodeMap(v reflect.Value) error {
	keys := v.MapKeys()
	sortedKeys := make([]string, len(keys))

	for i, key := range keys {
		switch key.Kind() {
		case reflect.String:
			sortedKeys[i] = key.String()
		case reflect.Interface:
			sortedKeys[i] = fmt.Sprint(key.Elem().Interface())
		default:
			sortedKeys[i] = fmt.Sprint(key.Interface())
		}
	}
	sort.Strings(sortedKeys)

	s.builder.WriteByte('{')
	for i, keyStr := range sortedKeys {
		if i > 0 {
			s.builder.WriteByte(',')
		}
		s.doEncodeString(keyStr)
		s.builder.WriteByte(':')

		var mapKey reflect.Value
		for _, key := range keys {
			var keyString string
			switch key.Kind() {
			case reflect.String:
				keyString = key.String()
			case reflect.Interface:
				keyString = fmt.Sprint(key.Elem().Interface())
			default:
				keyString = fmt.Sprint(key.Interface())
			}
			if keyString == keyStr {
				mapKey = key
				break
			}
		}

		if err := s.doEncode(v.MapIndex(mapKey)); err != nil {
			return err
		}
	}
	s.builder.WriteByte('}')
	return nil
}

func (s *OrderedSerializer) doEncodeStruct(v reflect.Value) error {
	t := v.Type()
	fields := make([]struct {
		name  string
		value reflect.Value
	}, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		name := field.Tag.Get("json")
		if name == "-" {
			continue
		}
		if comma := strings.Index(name, ","); comma != -1 {
			name = name[:comma]
		}
		if name == "" {
			name = field.Name
		}

		value := v.Field(i)
		if strings.Contains(field.Tag.Get("json"), ",omitempty") {
			if isEmptyValue(value) {
				continue
			}
		}

		fields = append(fields, struct {
			name  string
			value reflect.Value
		}{name, value})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].name < fields[j].name
	})

	s.builder.WriteByte('{')
	for i, field := range fields {
		if i > 0 {
			s.builder.WriteByte(',')
		}
		s.doEncodeString(field.name)
		s.builder.WriteByte(':')
		if err := s.doEncode(field.value); err != nil {
			return err
		}
	}
	s.builder.WriteByte('}')
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}
