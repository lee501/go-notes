package main

import "fmt"

/*
	考点：
		结构体属性为引用类型时， 函数参数为结构体值传递，更改属性的时候，引用类型值变了会改变最终值
*/

type T struct {
	ls []int
}

type V struct {
	age int
}
//结构体值为切片
func fo(t T) {
	t.ls[0] = 100
}
//结构体属性为普通
func vo(v *V) {
	v.age = 20
}
func main() {
	t := new(T)
	t.ls = []int{1,2,3}
	fo(*t)
	fmt.Println(t.ls[0])

	v := new(V)
	v.age = 10
	vo(v)
	fmt.Println(v.age)
}
