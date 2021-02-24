package main

import "fmt"

/*
	函数传参为切片，通过append对slice追加元素，不改变原来值的变化
	解决办法： 1. 通过返回值的方式处理
 			  2. 通过传递指针的方式处理
*/
func main() {
	s := make([]int, 3, 4)
	change(s)
	fmt.Println(s) // [0, 0, 0]
}

func change(s []int) {
	s = append(s, 5) //此处没有扩容地址还是main中s的地址
	fmt.Println(s) //内部值为[0 0 0 5]， 不改变main中的输出值
	s[0] = 1 // 未扩容的情况下，通过下标更改，外边值变为[1 0 0]
	s = append(s, 6) //此处扩容，地址改变
	s[1] = 2 // 不改变外部的值[1 0 0]
	fmt.Println(s) // [1, 2, 0, 5, 6]
}
//通过传递指针解决append slice issue
func usePoint(s *[]int) {
	*s = append(*s, 5)
	*s = append(*s, 6)
	fmt.Println(s) //[0 0 0 5 6] 外部的值也为这个
}
