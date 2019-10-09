package main

import "fmt"

/*
	考点： 关于引用类型的变量声明和赋值问题
		1 使用var m map[string]int， 这种只声明了变量，没有分配内存空间，所以不能进行赋值
		2 使用make或字面量来初始化map
*/

func main() {
	//错误示范
	//var i map[string]int
	//i["lee"] = 1
	//fmt.Println(i)

	//	正确初始化方法
	m := make(map[string]int)
	m["lee"] = 1

	if v, ok := m["lee"]; ok {
		fmt.Println(v)
	}
}
