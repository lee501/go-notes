package main

import (
	"fmt"
	"reflect"
)

type UserDemo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u UserDemo) RelectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func demoTest() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
	//DoFieldAndMethod(user)
}

// 参数对象为intereface{}
func DoFieldAndMethod(input interface{}) {
	// 获取类型
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())
	// 获取值
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		// 通过reflect.Type的Field获取其Field
		field := getType.Field(i)
		// 通过Field的Interface()得到对应的value
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 获取方法
	// 先获取interface的reflect.Type, 通过.NumMethod()进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
