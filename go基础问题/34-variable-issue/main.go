package main

import "fmt"

/*
	变量简短声明符号:=
			1. 作用域中， 导致变量隐藏问题
			2. 左边赋值多个变量时，只需保证一个变量是新声明的即可
*/
func main() {
	x := 1
	fmt.Println(x)
	//x变量隐藏
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x)
}
