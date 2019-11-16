package main

import "fmt"

/*
	interface内部结构包含静态类型动态类型以及动态值
		当且仅当动态值和动态类型为nil，接口类型值才为nil
		例子：var x *int = nil（动态类型为*int，值为nil）
*/

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non empty interface")
}
func main() {
	var x *int = nil
	//输出non empty
	Foo(x)
}
