package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string
	var vKind = str.Kind()

	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSetted, err := mapping(vPtr.Elem(), field, setter, tag)
		if err != nil {
			return false, err
		}
		if isNew && isSetted {
			value.Set(vPtr)
		}
		return isSetted, nil
	}
	s := &str
	Tee(*s)
}

func Tee(msg string) {
	if msg == "abc" {
		fmt.Println("ok")
	}
}
