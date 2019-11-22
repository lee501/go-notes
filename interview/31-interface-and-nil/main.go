package main

import "fmt"

/*
	类型断言语法： i.(Type)
		编译时会自动检测 i 的动态类型与 Type 是否一致
*/
func main() {
	x := interface{}(nil)
	fmt.Println("x:", x)

	y := (*int)(nil)
	fmt.Println(y)

	a := y == x
	b := y == nil
	_, c := x.(interface{})
	fmt.Println(a, b, c)

	var m interface{}
	fmt.Println(m)
}
