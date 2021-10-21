package main

import (
	"fmt"
	"sync"
)

func main() {
	//multi channel ctrl
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// 等待N个后台线程完成, 此处不能使用range来处理, fatal error: all goroutines are asleep - deadlock!
	for i := 0; i < cap(done); i++ {
		<-done
	}

	//waitgroup 控制
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}
