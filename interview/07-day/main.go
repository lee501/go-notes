package main

import "fmt"

/*
	append向切片追加元素， 是底层数组重新分配内存
	切片是引用类型，修改具体的元素，改变原切片
*/
func main() {
	i := []int{1,2,3}
	fmt.Printf("%#v\n", i)
	app(i)
	fmt.Printf("%#v\n", i)
	ap(i)
	fmt.Printf("%#v\n", i)
}
//append重新分配底层数组，i的地址不同了
func app(i []int) {
	fmt.Printf("append前地址%p\n", i)
	i = append(i, 4)
	fmt.Printf("append后地址%p\n", i)
}
//修改切片元素
func ap(i []int) {
	i[0] = 4
}
