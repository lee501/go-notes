package main

import "fmt"

/*
	考点：
	defer和panic
		defer先入后出(类似栈的数据结构)
		defer在panic前可以执行
*/
func main() {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("输入前")
	}()

	defer func() {
		fmt.Println("输入中")
	}()

	defer func() {
		fmt.Println("输入后")
	}()

	panic("检查到异常")
}
