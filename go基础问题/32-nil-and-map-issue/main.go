package main

import (
	"errors"
	"fmt"
)

/*
	知识点：
		1. cap函数适用于数组，数组指针，slice和channel， 不适用与map
		2. nil用于表示interface、函数、map、slice、 channel的零值， 不指定变量类型的时候，编译报错
*/
func main() {
	//	make 给map分配内存时，可以指定第二个参数，不过编译时会被忽略
	m := make(map[string]int, 2)
	fmt.Println(len(m))

	//nil值需要指定变量类型
	var x interface{} = nil
	fmt.Println(x == nil)

	//不能使用短变量声明设置结构体字段值
	var data info
	var err error
	data.result, err = work()
	fmt.Printf("info: %+v\n", data)
	fmt.Println(err)
}
type info struct {
	result int
}

func work() (int, error){
	return 2, errors.New("这是测试")
}
