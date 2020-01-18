package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	知识点
		结构体匿名继承时，如果时匿名指针情况，需要初始化，否者无法调用匿名组合的值方法
*/
type T struct {}

func (*T) foo() {
	fmt.Println("指针方法")
}

func (T) bar() {
	fmt.Println("值方法")
}

type S struct {
	*T
}

func test1() {
	s := S{} //T为nil
	fmt.Printf("%#v",s) //main.S{T:(*main.T)(nil)}指针方法
	//s.bar() 会panic（因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针）
	s.foo()
}
func main() {
	//test1()

	//测试指针接受者
	test2()
}

/*
	将 Mutex 作为匿名字段时，需要注意：
		1. 相关的方法必须使用指针接收者
		2. 或者通过嵌入 *Mutex 来避免复制的问题，但需要初始化
*/
type data struct {
	sync.Mutex
}

func (d *data) test(s string)  {
	d.Lock()
	defer d.Unlock()
	for i:=0;i<5 ;i++  {
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}

func test2() {
	var wg sync.WaitGroup
	wg.Add(2)

	d := new(data)

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

//嵌入*mutex来处理

type data1 struct {
	*sync.Mutex
}
func (d data1) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i:=0;i<5 ;i++  {
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}
func test3() {
	var wg sync.WaitGroup
	wg.Add(2)

	d := data1{new(sync.Mutex)} //初始化

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}