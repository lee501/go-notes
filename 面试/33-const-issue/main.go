package main

import "fmt"

/*
	常量在程序运行时，不会被修改，常量未使用是能编译通过的
*/

const (
	x uint16 = 120
	y // 常量未赋值，跟上面值一样
	s = "abc"
	z
)
func main() {
	const a = 123
	//量未使用是能编译通过的
	const b = 1.23
	fmt.Println(a)

	/*
		uint16 120
		string abc
	*/
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}
