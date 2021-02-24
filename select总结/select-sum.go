package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//for i:=0; i < 1000; i++ {
	//	randomSelect()
	//}

	selectBlock()
}

//1 select的随机性
func randomSelect() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case <-ch:
		fmt.Println("random 1")
	case <-ch:
		fmt.Println("random 2")
	}
}

//2select block情况， 主线程中，当channel阻塞的，没有default关键字的时候，deadlock!
func selectBlock() {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("random 1")
	case <-ch:
		fmt.Println("random 2")
	//default:
	//	//default防止panic
	//	fmt.Println("exit")
	}
}

//3超时机制
func selectTime() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("from ch")
	case <-timeout:
		fmt.Println("time out")
	case <-time.After(3 * time.Second):
		//使用time.After
		fmt.Println("time after")
	}
}

//4检查channel是否满了
func selectBuffer() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}

//5 for select用法
func selectFor() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	defer func() {
		close(ch)
	}()

	wg.Add(1)
	go func() {
		LOOP:
			for {
				select {
				case m := <-ch:
					fmt.Println(m)
					break LOOP
				}
			}
		fmt.Println("循环外")
		wg.Done()
	}()
	ch <- 1
	fmt.Println("主线程")
	wg.Wait()
}