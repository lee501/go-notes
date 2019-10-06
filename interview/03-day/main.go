package main

import "fmt"

/*
	make和new区别
		new(T)返回一个指针*T，该指针指向新分配的类型为T的零值，适用于值类型：数组和结构体
		make(T)返回的是初始化后的T的引用, 适用于slice map channel
		new分配的空间被清零，make分配后，会进行初始化
	声明指针不会开辟内存地址,只是准备要指向内存某个空间
*/
func main() {
	//声明值类型和引用类型区别
	//数组是值类型
	a := new([2]int)
	fmt.Println(a)
	//切片是引用类型, 不应该使用new
	b := new([]int)
	fmt.Println(b)
	//(*b)[0] = 1会报错， index out of range
}
