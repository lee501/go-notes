package main

import "fmt"

/*
	map初始化需要分配空间，、
	结构体作为map的值时，不能直接赋值结构体某个字段，拒绝直接寻址
*/

type P struct {
	name string
	sex string
}

func main() {
	m := make(map[string]P)
	//结构体不能直接赋值字段
	//m["a"].name = "lee"
	m["a"] = P{name: "lee", sex: "man"}
	fmt.Println(m)
	//当使用指针的时候可以直接赋值
	n := make(map[string]*P)
	p := &P{"lee", "man"}
	n["a"] = p
	n["a"].name = "test"
	fmt.Println(n["a"].name)
}
