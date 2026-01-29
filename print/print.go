package print

import (
	"fmt"
	"reflect"
)

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
}
