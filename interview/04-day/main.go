package main

import "fmt"

func main() {
	//考点：
	//不能使用new来初始化一个切片，append无法对*[]int类型的指针进行操作
	//list := new([]int)
	list := make([]int, 1)
	list = append(list, 1)
	fmt.Println(list) //输出[0,1]

	/*
		考点：多参数，append不能直接对切片进行操作, 需要使用多参操作符处理
	*/
	s := []int{1,2,3,4}
	n := []int{5,6}
	s = append(s, n...)
	fmt.Println(s)
}
