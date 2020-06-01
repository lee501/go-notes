package main

import (
	"fmt"
	"sync"
)

//func main() {
//	done := make(chan bool)
//	go func() {
//		fmt.Println("执行子线程")
//		done <- true
//	}()
//	<- done
//	fmt.Println("执行结束")
//}
func main() {
	var mutex sync.Mutex
	i := 6
	mutex.Lock()
	go func() {
		fmt.Println(i)
		mutex.Unlock()
	}()
	mutex.Lock()
}