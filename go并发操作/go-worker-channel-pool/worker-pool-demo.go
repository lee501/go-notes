package main

import (
	"fmt"
	"time"
)

//定义task类型，每个task抽象成一个函数
type Task struct {
	f func() error
}

//创建task
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

//执行task任务的方法
func (t *Task) Execute() {
	//调用任务绑定的函数
	t.f()
}

/*协程池的定义及操作*/
//定义Pool类型
type Pool struct {
	//接收task的入口
	EntryChannel chan *Task

	//协程池最大worker数量
	worker_num int

	//协程池内部任务就绪队列
	JobsChannel chan *Task
}

//创建一个协程池
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		worker_num: cap,
		JobsChannel: make(chan *Task),
	}
	return &p
}

//协程池创建worker执行工作
func (p *Pool) worker(workerId int)  {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("work ID", workerId, "执行完毕")
	}
}

//启动Pool
func (p *Pool) Run() {
	//根据worker num数量开启Worker, 每个worker是一个goroutine栽体
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i)
	}

	//由EntryChannel传递任务，送到JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	defer close(p.JobsChannel)
	defer close(p.EntryChannel)
}

//测试
func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now().Unix())
		return nil
	})

	p := NewPool(3)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.Run()
}