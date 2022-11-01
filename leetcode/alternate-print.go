package main

import (
	"fmt"
	"strings"
	"sync"
)

//交替打印算法
func alternatePrint() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Println(str[i : i+1])
				i++
				fmt.Println(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}

	}(&wait)
	number <- true
	wait.Wait()
}
