package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	msg := ""
	mu.Lock()
	go func() {
		msg = "你好"
		mu.Unlock()
	}()
	mu.Lock() //阻塞等待子线程先执行赋值
	fmt.Println(msg)
}
