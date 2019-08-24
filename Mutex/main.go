package main

import (
	"fmt"
	"sync"
)

/*
	互斥锁sync.Mutex使用场景
		多个goroutine访问同一个函数， 此函数操作一个全局变量
		使用go run -race查看竞争关系

	*读写锁提高了更高的并行性，也就是说允许多个线程几乎同一时间读取一个共享变量，而互斥锁不行
*/

var (
	//票数
	num = 100
	wg sync.WaitGroup
	mu sync.Mutex   //互斥锁
)
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go demo()
	}
	wg.Wait()
	fmt.Println(num, "程序执行结束")
}

func demo() {
	mu.Lock()
	for i := 0; i < 10; i++ {
		num = num - 1
	}
	mu.Unlock()
	wg.Done()
}