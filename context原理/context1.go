package main

import (
	"context"
	"fmt"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	//err := make(chan error)

	go func() {
		isSuccess := false
		if isSuccess {
			result <- 1}
		//} else {
		//	err <- errors.New("some error")
		//}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <- result:
		return nil
	}
}

func main() {
	//var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := Rpc(ctx, "http://rpc_1_url")
	//if err != nil {
	//	return
	//}
	fmt.Println(err)
	//rpc2
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	err := Rpc(ctx, "http://rpc_2_url")
	//	if err != nil {
	//		cancel()
	//	}
	//}()
	//
	//// RPC3调用
	//wg.Add(1)
	//go func(){
	//	defer wg.Done()
	//	err := Rpc(ctx, "http://rpc_3_url")
	//	if err != nil {
	//		cancel()
	//	}
	//}()
	//
	//// RPC4调用
	//wg.Add(1)
	//go func(){
	//	defer wg.Done()
	//	err := Rpc(ctx, "http://rpc_4_url")
	//	if err != nil {
	//		cancel()
	//	}
	//}()
	//
	//wg.Wait()
}