package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func demoReflectLearn() {
	t := &T{201, "mh103"}
	s := reflect.ValueOf(t)
	s = s.Elem()
	typeOfT := s.Type()
	fmt.Println(s.Type())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
