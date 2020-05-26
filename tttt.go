package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("进入goroutine")
		// 添加一个内容后控制台输出:1 true
		//ch<-1

		//关闭ch控制台输出:0 false
		//close(ch)
	}()
	c, d := <-ch
	fmt.Println(c, d)
	fmt.Println("程序执行结束")
}
