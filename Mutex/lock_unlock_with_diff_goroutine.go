package main

import (
	"fmt"
	"sync"
)

/*
	rwmutex 基于mutex，在mutex基础上增加了读写信号量
	读锁与读锁兼容，读与写互斥，写与写互斥
	* 无论是Mutex还是RWMutex都不会和goroutine进行关联，
	  这意味着它们的锁申请行为可以在一个goroutine中操作，释放锁行为可以在另一个goroutine中操作
*/
var mut sync.Mutex
var wi sync.WaitGroup
func main() {
	i := 2
	wi.Add(2)
	go change(&i)
	go read()
	wi.Wait()
	fmt.Println(i)
}

func change(i *int) {
	mut.Lock()
	*i = 5
	wi.Done()
}

func read() {
	mut.Unlock()
	wi.Done()
}