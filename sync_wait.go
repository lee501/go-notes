package main

import (
	"fmt"
	"sync"
)

func main() {
	r := in()
	fmt.Println(len(r))
}

func in() (res []int) {
	var wg sync.WaitGroup
	res = make([]int, 0)
	resChan := make(chan int, 1000)
	//defer close(resChan)
	run := func(resChan chan<- int, i int) {
		resChan <- i
		wg.Done()
	}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go run(resChan, i)
	}
	var wgConsume sync.WaitGroup
	wgConsume.Add(1)
	go func() {
		defer wgConsume.Done()
		//for {
		//	select {
		//	case r, ok := <-resChan:
		//		if !ok {
		//			return
		//		}
		//		res = append(res, r)
		//	}
		//}
		for r := range resChan {
			res = append(res, r)
		}
	}()

	wg.Wait()
	close(resChan)
	wgConsume.Wait()
	fmt.Println("all resChane consume finished")
	return res
}
