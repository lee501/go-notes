package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	1. sync.Pool 是goroutine并发安全
	2. 减少GC

	bufferpool := &sync.Pool {
		New: func() interface {} {
				println("Create new instance")
				return struct{}{}
			}
		}
	//申请对象
	buffer := bufferPool.Get()
	//释放对象(放回pool)
	bufferPool.Put(buffer)
*/

//统计实际的创建次数
var numCalCreated int32

func createBuffer() interface{} {
	atomic.AddInt32(&numCalCreated, 1)
	buffer := make([]byte, 512)
	return &buffer
}

func main() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}
	//并发go
	worker := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(worker)

	for i := 0; i < worker; i++ {
		go func() {
			defer wg.Done()
			buf := bufferPool.Get()
			_ = buf.(*[]byte)
			defer bufferPool.Put(buf)
		}()
	}

	wg.Wait()

	fmt.Printf("%d buffer obj was created.\n", numCalCreated)
}
