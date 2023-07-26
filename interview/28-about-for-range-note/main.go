package main

import "fmt"

/*
	range表达式是副本参与循环
		注意切片和数组： range会复制对象，而不是不是直接在原对象上操作
					   使用range迭代遍历引用类型时，底层的数据不会被复制
*/
func main() {
	////切片
	//testSlice()
	////数组
	//testArray()
	////数组地址传递给range
	//testArray1()
	testSliceAndAppend()
}

func testSlice() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	//r =  [1 12 13 4 5]
	// a =  [1 12 13 4 5]
}

func testArray() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	//r =  [1 2 3 4 5]
	//a =  [1 12 13 4 5]
}

func testArray1() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	//r =  [1 12 13 4 5]
	//a =  [1 12 13 4 5]
}

func testSliceAndAppend() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n%p\n%p\n", a, &a[2], &a[3])
	var r = make([]int, 0)
	//引用类型，a的长度不变
	for i, v := range a {
		if i == 1 {
			//注意原切片删除元素的坑
			a = append(a[:1], a[2:]...)
		}
		r = append(r, v)
	}
	fmt.Printf("%p\n%p\n%p\n", a, &a[2], &a[3])
	fmt.Println(a)
	fmt.Println(r)
	//[1 3 4 5]
	//[1 2 4 5 5]
}
