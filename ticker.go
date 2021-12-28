package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker := time.NewTicker(3 * time.Second)
	i := 0
	for {
		i += 3
		select {
		case <-ticker.C:
			fmt.Println("---------", i)
		}
		if i > 10 {
			break
		}
	}
	ticker.Stop()
	fmt.Println("ticker finish")
	
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	//if !timer.Stop() {
	//	<- timer.C
	//}
	//timer.Reset(2 * time.Second)
	for {
		i += 3
		select {
		case <-timer.C:
			timer.Reset(2 * time.Second)
			fmt.Println("---------", i)
		}
	}
}
