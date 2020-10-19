package lib

import (
	"reflect"
	"strconv"
)

func AutoSetLength(v interface{}, tag string) interface{} {
	rawType := reflect.TypeOf(v)
	var value reflect.Value
	isPrt := false
	switch rawType.Kind() {
	case reflect.Ptr:
		value = reflect.ValueOf(v)
		isPrt = true
		if rawType.Elem().Kind() != reflect.Struct {
			value.Elem().Set(reflect.ValueOf(AutoSetLength(value.Elem().Interface(), tag)))
			return value.Interface()
		}

	case reflect.Slice:
		value = reflect.ValueOf(v)
		for i := 0; i < value.Len(); i++ {
			value.Index(i).Set(reflect.ValueOf(AutoSetLength(value.Index(i).Interface(), tag)))
		}
		return value.Interface()
	case reflect.Map:
		value = reflect.ValueOf(v)
		mr := value.MapRange()
		if mr == nil {
			return v
		}
		for mr.Next() {
			value.SetMapIndex(mr.Key(), reflect.ValueOf(AutoSetLength(mr.Value().Interface(), tag)))
		}
		return value.Interface()
	case reflect.Array:
		value = reflect.New(reflect.TypeOf(v)).Elem()
		value.Set(reflect.ValueOf(v))
		for i := 0; i < value.Len(); i++ {
			value.Index(i).Set(reflect.ValueOf(AutoSetLength(value.Index(i).Interface(), tag)))
		}
		return value.Interface()
	case reflect.Struct:
		value = reflect.New(rawType)
		value.Elem().Set(reflect.ValueOf(v))
	default:
		return v
	}
	vtype := value.Elem().Type()
	for i := 0; i < vtype.NumField(); i++ {
		value.Elem().Field(i).Set(reflect.ValueOf(AutoSetLength(value.Elem().Field(i).Interface(), tag)))
		structInfo := value.Elem().Type().Field(i)
		if s := structInfo.Tag.Get("autoLen"); s != "" {
			var slen = 0
			var b bool

			if f, ok := vtype.FieldByName(s); ok {
				m := value.Elem().FieldByIndex(f.Index)
				if slen, b = autoSetLengthGetLen(m); !b {
					continue
				}
			}
			switch structInfo.Type.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				value.Elem().Field(i).SetInt(int64(slen))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value.Elem().Field(i).SetUint(uint64(slen))
			case reflect.String:
				value.Elem().Field(i).SetString(strconv.Itoa(slen))
			case reflect.Float32, reflect.Float64:
				value.Elem().Field(i).SetFloat(float64(slen))
			default:
				continue
			}
		}
	}
	if isPrt {
		return value.Interface()
	} else {
		return value.Elem().Interface()
	}
}
func autoSetLengthGetLen(value reflect.Value) (i int, b bool) {
	defer func() {
		if recover() != nil {
			b = false
		}
	}()
	if value.Kind() == reflect.Ptr {
		return autoSetLengthGetLen(value.Elem())
	}
	return value.Len(), true
}
