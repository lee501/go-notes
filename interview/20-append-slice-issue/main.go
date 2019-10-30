package main

import "fmt"
/*
	append slice issue： 扩容的时候重新分配底层数组
						 不扩容的时用的是原数组
*/
func main() {
	slice := make([]int, 5, 5)
	//扩容
	change(slice...)
	fmt.Println(slice)
	//不扩容
	change(slice[0:2]...)
	fmt.Println(slice)
}

func change(s ...int) {
	s = append(s, 3)
}