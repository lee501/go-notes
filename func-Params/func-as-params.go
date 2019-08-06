package main

import "fmt"

//函数是引用类型，传递的是地址
//函数作为参数的demo
func demo(iner func(name string)) {
	fmt.Println("demo started")
	iner("lee")
}

func main() {
	demo(func(name string){
		fmt.Println("name is ", name)
	})
}
