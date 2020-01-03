package main

import "fmt"

/*
	知识点：变量隐藏
		1. 变量使用简短号:=时， 左边出现多个变量，只需保证至少有一个变量是新声明的。
		2. 若出现作用域，会导致变量隐藏的问题

*/

func main() {
	x := 1
	//变量隐藏
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x)
}
