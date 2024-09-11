package main

import (
	"fmt"
	"time"
)

func main(){
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	select {
		case <- timeout:
			fmt.Println("exceed time")
	}
}
