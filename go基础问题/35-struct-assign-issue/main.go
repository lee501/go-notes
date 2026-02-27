package main

import "fmt"

/*
	结构体字段不可以用简短符号赋值
*/
type foo struct {
	bar int
}

func main() {
	var f foo
	f.bar = 1
	fmt.Println(f)

	m := new(foo)
	m.bar = 1
	fmt.Println(m.bar)

	BitRever()
}

/*
	有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数a取反，结果为 -(a+1) ）

	左移右移：  通过原码操作的
	位运算符：  使用的是补码
*/

func BitRever() {
	//00000011
	//11111100
	//11111101
	var a int8 = -3
	//00000010
	fmt.Println(^a)  // => 2
}
