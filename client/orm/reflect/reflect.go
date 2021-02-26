package reflect

import (
	"reflect"
	"strings"
)

// Field 字段
type Field struct {
	Name  string      `json:"name"`
	Op    string      `json:"op"`
	Value interface{} `json:"value"`
}

// Fields 获取字段信息
func Fields(condition interface{}) []*Field {
	t := reflect.TypeOf(condition)
	v := reflect.ValueOf(condition)
	fields := make([]*Field, 0, 8)
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			fields = resolveField(t.Field(i), v.Field(i), fields)
		}
	case reflect.Ptr:
		tl := t.Elem()
		vl := v.Elem()
		for i := 0; i < tl.NumField(); i++ {
			fields = resolveField(tl.Field(i), vl.Field(i), fields)
		}
	}
	return fields
}

func resolveField(f reflect.StructField, v reflect.Value, fields []*Field) []*Field {
	if tag, ok := f.Tag.Lookup("condition"); ok {
		if field := resolveTag(tag); field != nil {
			field.Value = v.Interface()
			fields = append(fields, field)
		}
	}
	return fields
}

func resolveTag(tag string) *Field {
	field := new(Field)
	pairs := strings.Split(tag, ";")
	for _, pair := range pairs {
		sl := strings.Split(pair, ":")
		if len(sl) == 2 {
			switch sl[0] {
			case "name":
				field.Name = sl[1]
			case "op":
				field.Op = sl[1]
			}
		}
	}
	if field.Name == "" || field.Op == "" {
		return nil
	}
	return field
}

// IsSlice 是否为slice
func IsSlice(v interface{}) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Slice
}

// IsNil 是否为nil
// chan, func, interface, map, pointer, or slice value
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	defer func() {
		recover()
	}()
	v := reflect.ValueOf(value)
	return v.IsNil()
}

// IsZero 是否为零值
func IsZero(value interface{}) bool {
	if value == nil {
		return true
	}
	defer func() {
		recover()
	}()
	v := reflect.ValueOf(value)
	return v.IsZero()
}
