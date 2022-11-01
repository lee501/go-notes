package main

import (
	"sync"
)

func sum(data []int) int {
	s := 0
	l := len(data)
	const N = 5
	seg := l / N

	var wg sync.WaitGroup
	wg.Add(N) // 直接加N个

	for i := 0; i < N; i++ {
		go func(i int) {
			tmpS := data[i*seg : (i+1)*seg]
			ll := len(tmpS)
			for j := 0; j < ll; j++ {
				s += tmpS[j]
			}
			wg.Done() // 一个goroutine运行完
		}(i)
	}
	wg.Wait() // 等N个goroutine都运行完

	return s
}

/*
func main() {
	data := []int{1,2,3,4,5,6,7,7,8,8,8,7,9,10}
	s := sum(data)
	fmt.Println(s)
}
*/
