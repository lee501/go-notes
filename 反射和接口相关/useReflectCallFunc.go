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

	//根据反射对象 NumIn 方法返回的参数个数创建 argv 数组
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Kind())
}

/*
type User struct {
	Id int
	Name string
}

func DoFieldMethod(input interface{}) {
	t := reflect.TypeOf(input)
	fmt.Println(t.Name())
	fmt.Println(t.NumField())

	v := reflect.ValueOf(input)
	fmt.Println(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(field).Interface()
		fmt.Println(field.Name, field.Type, value)
	}
}
*/
