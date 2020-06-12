package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sum int64
func worker_atomic(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var i int64
	for i = 0; i < 10000; i++ {
		atomic.AddInt64(&sum, i)
	}
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go worker_atomic(&wg)
	go worker_atomic(&wg)
	wg.Wait()
	fmt.Println("执行结束", sum)
}
