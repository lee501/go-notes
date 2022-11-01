package main

import (
	"fmt"
	"time"
)

func selectWait() {
	start := time.Now()
	wait := make(chan int, 1)
	go func() {
		fmt.Println("做点东西")
		time.Sleep(1 * time.Second)
		wait <- 2
	}()
	fmt.Println("这里是主程序")
	select {
	case nums := <-wait:
		fmt.Println(nums)
	case <-time.After(3 * time.Second):
		fmt.Println("3秒后")
	}
	fmt.Println(time.Since(start))
}
