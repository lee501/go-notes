package main

import "fmt"

//for range的特别之处:会一直阻塞当前协程，如果在其他协程中调用了close(ch),那么就会跳出for range循环
//demo
func forRange() {
	ch := make(chan string)
	ch2 := make(chan int)

	go func() {
		for i := 97; i < 123; i++ {
			ch <- fmt.Sprintf("%c", i)
		}
		//用来跳出range循环
		close(ch)
		ch2 <- 1
	}()

	go func() {
		for c := range ch {
			fmt.Println(c)
		}
		fmt.Println("check 阻塞")
	}()
	<-ch2
	fmt.Println("执行结束")
}
