package main

import (
	"fmt"
	"time"
)

func breakDemo() {

	ticker := time.NewTicker(time.Second * 2)
	timer := time.NewTimer(time.Second * 9)
cycle:
	for {
		select {
		case <-ticker.C:
			fmt.Println("*********ticker  break********")
		case <-timer.C:
			fmt.Println("break exit")
			break cycle
		}
	}

	ticker.Stop()
	timer.Stop()
	fmt.Println("break test over!")
}
