package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName string
	Name func(first, last string) string
}

func main() {
	person := &Person{
		FirstName: "lee",
		LastName: "andy",
		//结构体中包括匿名方法字段的初始化
		Name: func(first, last string) string {
			contactName := []string{first, last}
			return strings.Join(contactName, " ")
		},
	}
	fmt.Println(person.Name(person.FirstName, person.LastName))
}

/**
1。 //匿名结构体的赋值
m := struct {
		X, Y, Z float32
}{1, 2, 3}

2。 struct中，无论使用的是指针的方式声明还是普通方式，访问其成员都使用"."，
	在访问的时候编译器会自动把 stu2.name 转为 (*stu2).name。
*/

//工厂模式自定义构造函数
type Student struct {
	name string
	age int
	Class string
}

func NewStu(name, class string, age int) *Student {
	return &Student{name, age, class}
}

//tag
//tag可以为结构体的成员添加说明或者标签便于使用,这些说明可以通过反射获取到
type Student1 struct {
	Name string `json: "name"`
	Age int `json: "age"`
}
//json序列化
func ToJson() (string, error){
	stu := Student1{Name: "lee", Age: 22}
	data, err := json.Marshal(stu)
	if err != nil{
		fmt.Println("json encode failed err:",err)
		return "", err
	}
	return string(data), nil
}
