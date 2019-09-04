package main

import (
	"fmt"
	"sync"
)

//struct{}是一个无元素的结构类型, 大小为0， 通常作为占位符
//struct {} {}是一个复合字面量，它构造了一个struct {}类型结构值，该结构值也是空的：

type m chan struct{}

var wg sync.WaitGroup
func main() {
	n := make(m, 1)
	n <- struct {}{}
	select {
	case <- n:
		fmt.Println("空值")
	}

	test := make(chan int)
	go func() {
		test <- 1
	}()

	<- test
	fmt.Println("关闭")
}
