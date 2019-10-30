package main

import "fmt"

/*
	考点:
		for range循环,会创建每个元素的副本，所以v的地址是相同的
		range 表达式复制一个副本_, v := range slice， 这里的slice是副本引用
*/

func main() {
	slice := []int{1,2,3,4}
	//这里的v是循环对象每个元素的副本
	for _, v := range slice {
		fmt.Printf("%p\n", slice)
		fmt.Println(v)
	}

	m := make(map[int]*int)
	for key, val:= range slice {
		//将切片的每个地址作为map值，需要赋值一个新的变量
		value := val
		m[key] = &value
	}
	//此时得到的结果为0->1, 1->2
	for k, v := range m {
		fmt.Println(k, "----->", *v)
	}
	/*
	1
	0xc000094000
	2
	0xc000094000
	3
	0xc000094000
	4
	0xc000094000
	0 -----> 1
	1 -----> 2
	2 -----> 3
	3 -----> 4
	*/
}
