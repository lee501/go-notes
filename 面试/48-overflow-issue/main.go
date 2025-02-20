package main

import "fmt"

/*
	int8 取值范围 -128～127
*/

func main() {
	/*
		var x int8 = 128 // constant 128 overflows int8
		y := x/-1
		fmt.Println(y)
	*/
}

/*
	go 预定义的nil可以被覆盖
*/

func overNil() {
	nil := 123
	fmt.Println(nil)
	//var _ map[string]int = nil 这里nil被覆盖为int类型， 不能赋值给map
}
