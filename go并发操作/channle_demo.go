package main

/*
var wait sync.WaitGroup
var w sync.WaitGroup

func main() {
	ch := make(chan int)
	sign := make(chan int, 1)
	wait.Add(2)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		wait.Done()
	}(&wait)

	go func(wait *sync.WaitGroup) {
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		wait.Done()
	}(&wait)

	go func() {
		wait.Wait()
		sign <- 1
	}()
	w.Add(1)
	go func(wait *sync.WaitGroup) {
		sum := 0
	END:
		for {
			select {
			case m := <-ch:
				sum += m
				fmt.Println(sum)
			case <-sign:
				w.Done()
				break END
			}
		}
		fmt.Println("循环结束")
	}(&w)
	w.Wait()
	fmt.Println("执行结束")
}
*/
