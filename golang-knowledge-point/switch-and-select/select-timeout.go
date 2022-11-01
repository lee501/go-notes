package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		println("time out issue 1")
	//或使用time.After机制，回传chan time.Time
	case <-time.After(time.Second * 1):
		println("time out after")
	}

	//多个channel读取， 使用for + select机制实现
	i := 0
	ch0 := make(chan string)
	defer func() {
		close(ch0)
	}()

	go func() {
	LOOP:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++
			select {
			case m := <-ch:
				println(m)
				break LOOP
			default:
				println(i)
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ch0 <- "stop"
}
