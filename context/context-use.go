package main

import (
	"context"
	"fmt"
)

func testWorkContext() {
	ch := make(chan int)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer func() {
		fmt.Println("Main defer: cancel context")
		cancel()
		<-ch
	}()
	go workerContext(ctxWithCancel, ch)
}

func workerContext(ctx context.Context, ch chan int) {
	//ctxWithTime, cancel := context.WithTimeout(ctx, time.Duration(150) * time.Millisecond)
	//defer func() {
	//	fmt.Println("workerContext cancelled")
	//	cancel()
	//}()
	defer close(ch)
	fmt.Println("worker context start")
	select {
	case <-ctx.Done():
		fmt.Println("cancel infer woker to close")
	default:
		ch <- 1
	}
}
