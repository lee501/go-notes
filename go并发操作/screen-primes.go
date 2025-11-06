package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNatural(ctx)

	for i := 0; i< 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		//过滤生成新的channel
		ch = PrimeFilter(ctx, ch, prime)
	}
	cancel()
}

func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:

			}
		}
	}()
	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i % prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:

				}
			}
		}
	}()
	return out
}