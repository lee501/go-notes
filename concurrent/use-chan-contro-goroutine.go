package main

import (
	"fmt"
	"sync"
)

//func main() {

//}
func main() {
	//使用mutex lock控制并发
	var mutex sync.Mutex
	i := 6
	mutex.Lock()
	go func() {
		fmt.Println(i)
		mutex.Unlock()
	}()
	mutex.Lock()
	//使用chan控制并发
	done := make(chan bool, 1)
	go func() {
		fmt.Println("执行子线程")
		done <- true
	}()
	<-done
	fmt.Println("执行结束")
}
