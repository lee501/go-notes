package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var msignal chan os.Signal

func dealSignal() {
	for {
		<-msignal
		fmt.Println("Get Signal ...")
	}
}

func main() {
	i := 0
	msignal = make(chan os.Signal, 1)
	signal.Notify(msignal, os.Interrupt)
	go dealSignal()
	for {
		i++
		time.Sleep(1 * time.Second)
		if i > 5 {
			break
		}
	}
}
