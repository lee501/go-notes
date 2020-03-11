package main

import (
	"fmt"
	"time"
)

//多个channel读取， 使用for + select机制实现
func main() {
	i := 0
	ch := make(chan string)
	defer func() {
		close(ch)
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
	ch <- "stop"
}
