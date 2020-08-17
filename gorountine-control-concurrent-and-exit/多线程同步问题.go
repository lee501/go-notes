package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//使用channel和WaitGroup同步
func main() {
	done:=make(chan int ,100)
	defer close(done)
	//开启线程
	for i := 1; i <= cap(done); i++ {
		go func(i int) {
			fmt.Println("开启线程", i)
			done <- i
		}(i)
	}
	//使用channel阻塞的方式来出来同步
	//此处不能使用range 会引起主线程deadline
	for i := 0; i< 100; i++ {
		m := <-done
		fmt.Println(m, "线程关闭")
	}
 	fmt.Println("执行完毕")

	//2. 使用waitGroup方式防同步协程
	var wg sync.WaitGroup
	for i:=0; i<100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i, "开始执行")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//使用context来控制
func foo(ctx context.Context, name string) {
	go bar(ctx, name)
	for  {
		select {
		case <- ctx.Done():
			fmt.Println(name, "foo exit")
			return
		case <- time.After(2 * time.Second):
			fmt.Println("foo 超时")
		}
	}
}

func bar(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B Exit")
			return
		case <-time.After(2 * time.Second):
			fmt.Println(name, "B do something")
		}
	}
}

func Test() {
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, "foobar")
	fmt.Println("client release connection, need to notify A, B exit")
	time.Sleep(5 * time.Second)
	cancel() //mock client exit, and pass the signal, ctx.Done() gets the signal  time.Sleep(3 * time.Second)
	time.Sleep(3 * time.Second)
}