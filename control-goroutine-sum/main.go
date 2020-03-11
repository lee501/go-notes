package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	jobs := 10
	//控制线程数量
	pool := 3
	jobChans := make(chan int, pool)
	for i := 0; i < pool; i++ {
		go func() {
			for ch := range jobChans {
				fmt.Println("hello" + strconv.Itoa(ch))
				wg.Done()
			}
		}()
	}
	for i := 0; i < jobs; i++ {
		wg.Add(1)
		jobChans <- i
		fmt.Printf("index: %d, goroutine number: %d\n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("done")
}
