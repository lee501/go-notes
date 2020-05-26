package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func main() {
	ch := make(chan int)
	wait.Add(3)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 5; i++{
			ch <- i
		}
		wait.Done()
	}(&wait)

	go func(wait *sync.WaitGroup) {
		for i := 0; i < 5; i++{
			ch <- i
		}
		wait.Done()
	}(&wait)

	go func(wait *sync.WaitGroup) {
		sum := 0
		END:

		for  {
			select {
			case m := <-ch:
				fmt.Println(m)
				sum += m
			default:
				fmt.Println("test")
				break END
			}
		}
		fmt.Println(sum)
		wait.Done()
		fmt.Println("循环结束")
	}(&wait)
	wait.Wait()
	fmt.Println("执行结束")
}

