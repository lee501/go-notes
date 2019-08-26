package main

import (
	"fmt"
	"time"
)

//主线程中，channel阻塞：产生死锁fatal error : all goroutines are asleep -deadlock!
func main()  {
	ch := make(chan int, 10)
	go func() {
		for i:=0; i<10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := range ch{
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second)
}
