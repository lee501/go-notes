package main

import (
	"fmt"
	"reflect"
)

func Add(a,b int) int {
	return a + b
}

func main() {
	v := reflect.ValueOf(Add)
	fmt.Println(v.Kind())
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	fmt.Println(t)
}
