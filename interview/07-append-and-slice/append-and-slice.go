package main

import "fmt"

/*
	append向切片追加元素，只有cap不够的情况下，底层数组才重新分配内存
	在make初始化切片的时候给出了足够的容量，append操作不会创建新的切片
	切片是引用类型，修改具体的元素，改变原切片

slice的实现就是复制时,
如果不超过容量那么底层数组数据共享, 但是长度数据或者说有效位数是不共享的, 各玩各的.

*/
func main() {
	i := []int{1, 2, 3}
	fmt.Printf("append前地址%p\n", i)
	i = append(i, 4)
	fmt.Printf("append前地址%p\n", i)
	//app(i)
	fmt.Printf("%#v\n", i)
	ap(i)
	fmt.Printf("%#v\n", i)
	s := []int{1, 2, 3, 4}
	//fmt.Printf("append前地址%p\n", s)
	//s = s[:2 + copy(s[2:], s[3:])]
	//fmt.Printf("append前地址%p\n", s)
	//fmt.Println(s)
	b := make([]int, len(s))
	copy(b, s)
	fmt.Println(b)
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
