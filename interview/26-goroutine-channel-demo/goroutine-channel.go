package main

import (
	"fmt"
	"time"
)

/*
	协程和channel关闭的demo
*/
func main() {
	ch := make(chan int, 20)
	//A
	go func() {
		for i := 0; i < 10; i++{
			ch <- i
		}
		//close(ch)
	}()
	//B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a:", a)
		}
	}()
	defer close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 5)
}
