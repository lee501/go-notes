package main

import "fmt"

/*
	知识点：
		当目标方法的接受者为指针类型时，被复制的是指针地址
*/

type N int

func (n *N) test() {
	fmt.Println(*n)
}
func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)
	f1() //13
	f2() //13
	//声明s切片，未初始化，值为nil，不能直接使用下标赋值
	var s []int
	fmt.Println(s)
}
