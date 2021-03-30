package main

import "fmt"

func Producer(t int, out chan <- int) {
	for i:=0; ; i++ {
		out <- i*t
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}