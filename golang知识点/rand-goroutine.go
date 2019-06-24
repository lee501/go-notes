package main

import (
	"fmt"
	"runtime"
	"sync"
)

//gorountine 的随机性和闭包
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	//外部for循环的变量i, 地址不变化, 最后i=10， go引用的是外部变量i
	for i := 0; i< 10; i++ {
		go func(){
			fmt.Println("A:", i)
			wg.Done()
		}()
	}
	//i为函数参数，函数值拷贝后，go func内部指向值拷贝地址
	for i := 0; i< 10; i++ {
		go func(i int){
			fmt.Println("B:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
