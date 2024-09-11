package main

import "fmt"

/*
	Golang中只有三种引用类型：slice、map、channel
	在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
	引用类型理解为指针：堆存储， 引用的是地址

	值类型：值的拷贝， 值存储在内存的栈中
*/


//值类型和引用类型的区别，用数组和slice例子
func main()  {
	//值类型
	a := [5]int{1,2,3,4,5}
	b := a
	c := new([3]int)
	b[2] = 6
	fmt.Println(a, b, cap(c))

	//引用类型
	d := make([]int, 5)
	d = []int{1,2,3,4}
	fmt.Println(d, len(d), cap(d))

	e := new([]int) // 返回指针&T{}
	fmt.Println(len(*e))
}