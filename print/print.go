package print

import (
	"fmt"
	"reflect"
	"strings"
)

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
}

func PrintSliceStruct(slice interface{}) {
	v := reflect.ValueOf(slice)

	if v.Kind() != reflect.Slice {
		fmt.Println("type is not slice")
		return
	}
	if v.Len() == 0 {
		fmt.Println("empty slice")
		return
	}

	elemType := v.Index(0).Type()
	if elemType.Kind() != reflect.Struct {
		fmt.Println("it is not slice of struct")
		return
	}

	cols := elemType.NumField()
	widths := make([]int, cols)

	for i := 0; i < cols; i++ {
		widths[i] = len(elemType.Field(i).Name)
	}

	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		for j := 0; j < cols; j++ {
			val := fmt.Sprintf("%v", row.Field(j).Interface())
			if len(val) > widths[j] {
				widths[j] = len(val)
			}
		}
	}

	printLine := func() {
		fmt.Print("+")
		for i := 0; i < cols; i++ {
			fmt.Print(strings.Repeat("-", widths[i]+2))
			fmt.Print("+")
		}
		fmt.Println()
	}

	printLine()

	fmt.Print("|")
	for i := 0; i < cols; i++ {
		fmt.Printf(" %-*s |", widths[i], elemType.Field(i).Name)
	}
	fmt.Println()

	printLine()

	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		fmt.Print("|")
		for j := 0; j < cols; j++ {
			fmt.Printf(" %-*v |", widths[j], row.Field(j).Interface())
		}
		fmt.Println()
	}

	printLine()
}
