package main

import (
	"fmt"
	"strconv"
	"sync"
)

//RWMutex可以添加多个读锁或一个写锁
//读写锁在锁的范围内对数据的操作
//互斥锁代码只能一个goroutine运行
func main() {
	var rwm sync.RWMutex
	//var r sync.Mutex
	var m = make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			rwm.Lock()
			fmt.Println(m)
			m["key" + strconv.Itoa(i)] = "value" + strconv.Itoa(i)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("执行结束")
}
