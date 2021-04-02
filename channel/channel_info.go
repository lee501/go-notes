package main

import (
	"fmt"
	"time"
)

func Process1(ch chan<- int) {
	defer close(ch)
	fmt.Println(1)
	ch <- 1
}

func Process2(ch <- chan int) {
	if i, ok := <-ch; ok {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go Process1(ch)
	go Process2(ch)
	time.Sleep(5*time.Second)
}
