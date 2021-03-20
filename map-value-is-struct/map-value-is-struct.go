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
	//直接初始化
	t := map[string]P{"key": P{name: "lee", sex: "nan"}}
	//赋值value 为temp, temp可寻址，此时可以更改temp的属性name和sex
	temp := t["key"]
	temp.name = "anne"
	//更改后的t中value不变， （struct为值类型）
	fmt.Println(t, temp)
	//当使用指针的时候可以直接赋值
	n := make(map[string]*P)
	p := &P{"lee", "man"}
	n["a"] = p
	n["a"].name = "test"
	fmt.Println(n["a"])
}
