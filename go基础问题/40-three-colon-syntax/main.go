package main

import "fmt"

/*
	...语法：
		1. 用于函数多个不定参数的情况
		2. 将slice打散进入传递
*/

func main() {
	var strs = []string{
		"test1",
		"test2",
		"test3",
	}
	importSlice(strs...)
}

func importSlice(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}
}
