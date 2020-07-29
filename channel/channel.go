package main

import (
	"fmt"
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

		for i := range ch{
			fmt.Println(i)
		}

	//done := make(chan bool) // 无缓冲通道
	//defer close(done)
	//
	//go func() {
	//	time.Sleep(9 * time.Second)
	//	fmt.Println("one done")
	//	done <- true
	//}()
	//
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	fmt.Println("two done")
	//	done <- true
	//}()
	//
	//// wait until both are done
	//for c := 0; c < 2; c++ {
	//	<-done
	//}
	//fmt.Println("handle1 done")
}

