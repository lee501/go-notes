package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	val := `{"name": "admin"}`
	type Role struct {
		Name string
		IsOn bool
	}
	var role Role
	json.Unmarshal([]byte(val), &role)
	fmt.Println(role)
	//r := in()
	//fmt.Println(len(r))
}

func in() (res []int) {
	var wg sync.WaitGroup
	res = make([]int, 0)
	resChan := make(chan int, 1000)
	//done := make(chan struct{})
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
		//defer func() {
		//	close(resChan)
		//	close(done)
		//}()
		//for {
		//	select {
		//	case r := <-resChan:
		//		res = append(res, r)
		//	case <-done:
		//		return
		//	}
		//}
		for r := range resChan {
			res = append(res, r)
		}
	}()

	wg.Wait()
	close(resChan)
	wgConsume.Wait()
	//done <- struct{}{}
	fmt.Println("all resChane consume finished")
	return res
}
