package main

import "time"

func main() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	ch := make(chan int)
	select {
	case <- ch:
	case <-timeout:
		println("time out issue 1")
	//或使用time.After机制，回传chan time.Time
	case <-time.After(time.Second * 1):
		println("time out after")
	}
}
