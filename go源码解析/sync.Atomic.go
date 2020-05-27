package main

/*
	sync/atomic
 		实现原理
			大致是向CPU发送对某一个块内存的LOCK信号，然后就将此内存块加锁，从而保证了内存块操作的原子性
		1. 可以在并发的场景下对变量进行非侵入式的操作，保证并发安全
		2. 解决的典型问题就是 i++和CAS（Compare-and-Swap）的线程安全问题
*/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	test1()
	//test2()
}

//count++ 并发不安全
func test1() {
	var wg sync.WaitGroup
	count := 0
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("test1花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}

//使用atomic.AddInt64(&count, 1)原子操作
func test2() {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("test2 花费时间：%d, count的值为：%d \n",time.Now().Sub(t),count)
}