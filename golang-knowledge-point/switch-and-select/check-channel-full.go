package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel is blocking")
	}
}
