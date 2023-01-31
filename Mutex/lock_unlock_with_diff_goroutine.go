package main

import (
	"fmt"
	"strconv"
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

func mutux() {
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

//RWMutex可以添加多个读锁或一个写锁
//读写锁在锁的范围内对数据的操作
//互斥锁代码只能一个goroutine运行
func rwmutex() {
	var rwm sync.RWMutex
	//var r sync.Mutex
	var m = make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			rwm.Lock()
			fmt.Println(m)
			m["key"+strconv.Itoa(i)] = "value" + strconv.Itoa(i)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("执行结束")
}
