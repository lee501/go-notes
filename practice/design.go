package main

import (
	"fmt"
	"sync"
)

//策略模式： 封装方法
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator
}

type Addition struct {

}

func (a *Addition) Apply(i, j int) int {
	return i + j
}

func main() {
	var operator Operator = &Addition{}
	op := Operation{operator}
	fmt.Println(op.Apply(1, 1))
}

//对象池
type Object struct {
	Name string
}
type Pool chan *Object

func NewPools(size int) *Pool {
	pool := make(Pool, size)
	defer func() {
		close(pool)
	}()
	for i:=0; i<size; i++ {
		pool <- new(Object)
	}
	return &pool
}

//单类模式
type singleton map[string]string
var (
	once sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

//使用channel和goroutine设计生成器
func Generate(first, end int) <-chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := first; i < end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	return ch
}

//原型
type Demo struct {
	Des string
}

func (d *Demo) Clone() *Demo  {
	r := *d
	return &r
}

func NewDemo(des string) *Demo {
	return &Demo{
		Des: "demo",
	}
}

