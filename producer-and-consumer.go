package main

import (
	"fmt"
	"time"
)
//声明任务channel以及生成者和消费者数量
var (
	taskChannel chan int
	m int
	n int
)

//定义消费者
type Consumer struct {
	buffer chan int  //消费者缓存
	num    int       //消费者id
}

func NewConsumer(id, c int) *Consumer {
	return &Consumer{
		buffer: make(chan int, c),
		num: id,
	}
}
//消费者缓存
func (c *Consumer) In(t int){
	c.buffer <- t
}

func (c *Consumer) Execute() {
	fmt.Printf("consumer %d run task %d\n", c.num, <-c.buffer)
}

//func main() {
//	m = 3
//	n = 2
//	taskChannel = make(chan int, m)
//	var pause chan struct{}
//	//根据m生成producer
//	for i:=1; i<=m; i++ {
//		go producer(taskChannel, i)
//	}
//
//	//缓存容量
//	c := m/n
//	//生成n个消费者
//	for i:=1; i<=n; i++ {
//		go consumer(taskChannel, i, c)
//	}
//	pause <- struct{}{}
//}

func producer(taskChan chan int, id int) {
	for {
		taskChan <- id
		time.Sleep(3*time.Second)
	}
}

func consumer(taskChan chan int, id, c int) {
	con := NewConsumer(id, c)
	for {
		for t := range taskChan{
			if len(con.buffer) < c {
				con.buffer <- t
			}
			con.Execute()
		}
	}
}
