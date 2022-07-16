package logger

import (
	"github.com/Squirrel-Network/gobotapi/utils/ordered_map"
	"reflect"
)

func internalSerialize(v any) any {
	if v != nil {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Struct:
			r := ordered_map.New()
			name := reflect.TypeOf(v).Name()
			if name != "" {
				r.Set("_", reflect.TypeOf(v).String())
			}
			for i := 0; i < reflect.ValueOf(v).NumField(); i++ {
				r.Set(reflect.TypeOf(v).Field(i).Name, internalSerialize(reflect.ValueOf(v).Field(i).Interface()))
			}
			return r
		case reflect.Ptr:
			if !reflect.ValueOf(v).IsNil() {
				return internalSerialize(reflect.ValueOf(v).Elem().Interface())
			}
		case reflect.Slice:
			var r []any
			tmpSlice := reflect.ValueOf(v)
			for i := 0; i < tmpSlice.Len(); i++ {
				r = append(r, internalSerialize(tmpSlice.Index(i).Interface()))
			}
			return r
		}
	}
	return v
}
