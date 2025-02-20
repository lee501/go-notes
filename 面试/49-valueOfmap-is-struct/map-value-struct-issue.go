package main

import "fmt"

type People struct {
	Name string
}

//map的value是struct, 注意不可寻址的问题
//解决1： 结构体使用指针， 解决2：赋值给另一个变量temp, 修改temp.Name, 再将temp赋值给map[key]
func main() {
	m := map[string]People{"key": {Name: "lee"}}
	//不可以使用m["key"].Name = "another"来赋值，
	temp := m["key"]
	temp.Name = "anne"
	m["key"] = temp
	fmt.Println(m["key"])

	//或者
	n := map[string]*People{"key": &People{Name: "lee"}}
	n["key"].Name = "llll"
	fmt.Println(n["key"])
}
