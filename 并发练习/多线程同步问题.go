package main

import (
	"fmt"
	"sync"
)

func main() {
	done:=make(chan int ,100)
	defer close(done)
	//开启线程
	for i := 1; i <= cap(done); i++ {
		go func(i int) {
			fmt.Println("开启线程", i)
			done <- i
		}(i)
	}
	//使用channel阻塞的方式来出来同步
	for i := 0; i< 100; i++ {
		m := <-done
		fmt.Println(m, "线程关闭")
	}
 	fmt.Println("执行完毕")

	var wg sync.WaitGroup
	for i:=0; i<100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i, "开始执行")
			wg.Done()
		}(i)
	}
	wg.Wait()
}
