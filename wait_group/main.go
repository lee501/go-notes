package main

import (
	"fmt"
	"sync"
)

//使用waitGroup来处理程序阻塞
/*
	type WaitGroup struct {
		noCopy noCopy
		state1 [12]byte
		sema	uint32
	}
	*method
		func (wg *WaitGroup) Add(delta int)
		func (wg *WaitGroup) Done()
		func (wg *WaitGroup) Wait()
*/
var wg sync.WaitGroup

func main() {
	wg.Add(5)
	for i := 1; i < 6; i++ {
		go func(i int) {
			fmt.Printf("第%d次执行\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
