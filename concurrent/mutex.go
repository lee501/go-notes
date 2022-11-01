package main

import (
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wt *sync.WaitGroup) {
	defer func() {
		wt.Done()
	}()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

/*

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Printf("执行结束，value为%v\n", total.value)
}
*/
