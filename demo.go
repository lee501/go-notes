package main

import (
	"fmt"
	"strings"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Array 实现Interface接口
type Array []int

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 匿名接口(anonymous interface)
type reverse struct {
	Interface
}

// 重写(override)
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// 构造reverse Interface
func Reverse(data Interface) Interface {
	return &reverse{data}
}

//func main() {
//	arr := Array{1, 2, 3}
//	rarr := Reverse(arr)
//	fmt.Println(arr.Less(0,1))
//	fmt.Println(rarr.Less(0,1))
//
//	var t Base
//	t = &Test{}
//	fmt.Println(t.Name())
//
//}

type Test struct {
	base
}

type base struct {
}

func (b *base) Name() string {
	return "base"
}

type Base interface {
	Name() string
}

//func main() {
//	var dict map[string]map[string]int64
//	dict = make(map[string]map[string]int64)
//	addItem(dict)
//	fmt.Println(dict["table"])
//}

func addItem(dc map[string]map[string]int64) {
	if dc == nil {
		fmt.Println("no gc")
	}
	fmt.Println("add item")
	dc["table"] = map[string]int64{"a": 1}
}

func main() {
	str := "12"
	for _, item := range strings.Split(str, ",") {
		fmt.Println(item)
	}
	fmt.Println(strings.Split(str, ",")[0])
	fmt.Println(32 << 23)
}
