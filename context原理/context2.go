package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	root := context.Background()
	//ctx1, cancel := context.WithDeadline(root, time.Now().Add(6 * time.Second))
	ctx1, cancel := context.WithCancel(root)
	ctx2 := context.WithValue(ctx1, "key2", "value2")
	wg.Add(1)
	go watch(&wg, ctx2)
	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止")
	cancel()
wg.Wait()
}

func watch(wg *sync.WaitGroup, ctx context.Context) {
	OUT:
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Value("key2"), "监控退出了。")
			break OUT
		default:
			fmt.Println(ctx.Value("key2"), "go rountine 监控中。。。")
			time.Sleep(2 * time.Second)
		}
	}
	wg.Done()
}