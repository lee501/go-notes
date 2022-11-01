package main

import (
	"context"
	"time"
)

/*
func main() {
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")
	//返回一个context，设置超时时间为3秒
	//ctxb, _ := context.WithTimeout(ctxa, time.Second*3)
	//go work(ctxb, "work2")
	////当work2超时关闭的时候，通知work3 quit
	//ctxc := context.WithValue(ctxb, "key", "custom value")
	//go workWithValue(ctxc, "work3")
	//time.Sleep(5*time.Second)
	//手动关闭work1
	cancel()
	time.Sleep(time.Second)
}
*/

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			println(name, " is running")
			time.Sleep(time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " quit")
			return
		default:
			value := ctx.Value("key").(string)
			println(name, " is running with value ", value)
			time.Sleep(time.Second)
		}
	}
}
