package main

import "fmt"

/*
	常量考点
		1. 常量在程序运行时不会被修改的量，未使用也可以编译通过
		2. 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同

	string考点
		string不能分配nil， 若空值使用""
*/

const (
	x uint16 = 120
	y
	s string = "abc"
	z
)
func main() {
	const a = 123
	const b = 1.23
	fmt.Println(a)

	fmt.Printf("%T %v\n", y, y)
}

