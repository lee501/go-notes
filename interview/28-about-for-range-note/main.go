package main

import "fmt"

/*
	range表达式是副本参与循环
		注意切片和数组： 表达式为数组，传递值副本
					   表达式为切片， 传递引用地址
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
	var a = [5]int{1,2,3,4,5}
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
	var a = [5]int{1,2,3,4,5}
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
	var a = []int{1,2,3,4,5}
	fmt.Printf("%p\n", a)
	var r = make([]int, 0)
	//此处a为原切片的引用
	for i, v := range a {
		if i == 0 {
			//此处a重新分配了地址
			a = append(a, 6, 7)
			fmt.Printf("%p\n", a)
		}
		r = append(r, v)
	}
	fmt.Println(r)
	fmt.Println(a)
}