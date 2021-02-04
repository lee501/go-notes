package main

import (
	"fmt"
	"sync"
)

func main() {
	//product chan size
	producer := &Producer{make(chan *Workload, 100)}
	Question2(producer)
}

// IWorkload 请勿修改接口
type IWorkload interface {
	// Process内包含一些耗时的处理，可能是密集计算或者外部IO
	Process()
}

// IProducer 请勿修改接口
type IProducer interface {
	// Produce每次调用会返回一个IWorkload实例
	// 当返回nil时表示已经生产完毕
	Produce() IWorkload
}

// 问题2：请编写函数Question2的实现如下功能
// 该函数输入一个IProducer实例，请调用Produce()方法生产IWorkload实例
// 对于每个IWorkload实例，调用其Process()方法执行其业务功能
// 重复调用Produce()直到生产完毕，对每个IWorkload进行Process()，
// 所有工作完成后Question2函数返回
//
// 要求：系统允许最大5个IWorkload同时进行Process()操作
//      请利用golang的并发特性完成此任务，单个并发的实现将不得分
//
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

func Question2(producer IProducer) {

	// ====== 在这里书写代码 ====== //
	if prod, ok := producer.(*Producer); ok {
		go func() {
			for i := 0; i< 100; i++ {
				prod.Produce()
			}
			close(prod.W)
		}()
		pool := 5 //控制process数量
		var wait sync.WaitGroup
		wait.Add(5)
		for i:= 1; i <= pool; i++ {
			go func(i int) {
				fmt.Println("线程", i)
				for w := range prod.W {
					w.Process()
				}
				wait.Done()
			}(i)
		}
		wait.Wait()
	}
}

type Producer struct {
	W chan *Workload
}

func (p *Producer) Produce() IWorkload {
	w :=  &Workload{}
	p.W <- w
	return w
}

type Workload struct {

}

func (w *Workload) Process() {
	println("w")
}

