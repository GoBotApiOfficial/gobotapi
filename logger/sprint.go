package logger

import (
	"fmt"
	"reflect"
)

func sprint(a ...any) string {
	for t := range a {
		if a[t] != nil {
			rV := reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Ptr {
				a[t] = reflect.ValueOf(a[t]).Elem().Interface()
			}
			rV = reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Struct || rV == reflect.Map {
				a[t], _ = Serialize(a[t])
			}
		}
	}
	r := fmt.Sprintln(a...)
	r = r[:len(r)-1]
	return r
}
