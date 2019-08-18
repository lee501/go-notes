package main

import "fmt"

//recover通常是用在defer内部，因为panic的时候，defer是执行的
//函数调用过程中， recover只恢复当前函数中的panic，恢复后当前函数结束，当前函数panic后部分不执行
//panic会一直向上传递的，碰到recover的时候，panic停止向上传递
func demo() {
	fmt.Println("demo上半部分")
	demo1()
	fmt.Println("demo下半部分")
}

//在demo1中添加recover， demo2的panic上传到demo1中，所以demo1下半部分不执行
func demo1() {
	defer func() {
		recover()
	}()
	fmt.Println("demo1上半部分")
	demo2()
	fmt.Println("demo1下半部分")
}

func demo2() {
	fmt.Println("demo2上半部分")
	panic("执行panic")
	fmt.Println("demo2下半部分")
}

func main() {
	demo()
}