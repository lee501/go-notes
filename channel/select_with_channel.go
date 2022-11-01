package main

import "fmt"

//for循环中使用select： 使用break标签退出循环
func forSelect() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
EXIT:
	for {
		select {
		case c := <-ch:
			fmt.Println(c)
		default:
			break EXIT
		}
	}
	fmt.Println("程序结束")
}
