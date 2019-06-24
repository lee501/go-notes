package main

import (
	"fmt"
	"sync"
)

//单项channel的初识化
//'''
//ch := make(chan int)
//ch1 := <-chan int(ch) //类型转换为单向读取channel
//ch2 := chan<- int(ch) //类型转换为单向写入channel
//'''
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() chan interface{} {
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- value
			println("Iter", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	th := threadSafeSet{s: []interface{}{1,2,3}}
	v := <-th.Iter()
	fmt.Println(v)
}
