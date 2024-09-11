package main

import (
	"fmt"
)

//退出协程的三种方式
/*
	1.使用for-range退出
		range可以感知channel的关闭，当channel被发送数据的协程关闭时候，range会结束，退出for循环
*/

func someDemo() {
	var ch = make(chan int)
	var ch1 = make(chan int)
	/*
		使用for-range退出
			使用场景，当协程只从1个channel读取数据的时候，range处理channel, 当channel关闭时， 协程自动退出
	*/
	go func(in <-chan int) {
		for v := range in {
			//当in关闭时候，输出0
			fmt.Println(v)
		}
	}(ch)

	/*
		2. for-select
		使用ok退出：解决的是当读channel关闭时，没有数据读取的程序正常关闭
			for-select可以持续处理多个channel, select提供多路复用的功能，不会在nil的通道上进行等待。  但是无法感知channel关闭,
			引发以下两个问题
			1.继续在关闭的channel上读，会读到0值或nil
			2.继续在关闭的通道上写，会panic
	*/
	//某个通道关闭后，需要退出协程， 直接使用return
	go func() {
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			default:

			}
		}
	}()

	/*
		select 不会在nil的通道上进行等待
		应用场景： 某个通道关闭了，不再处理该通道，继续处理其他的case
	*/
	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					ch = nil
				}
			case _, ok := <-ch1:
				if !ok {
					ch1 = nil
				}
			default:

			}
			//通道都关闭的时候退出循环
			if ch == nil && ch1 == nil {
				return
			}
		}
	}()
}

/*
	3. 使用退出通道退出
			应用场景：
			a. 接收协程退出，不告知发送协程，引发阻塞
			b. 启用工作协程处理任务，如何退出
*/
func worker(stopCh <-chan struct{}) {
	defer fmt.Println("worker exit")
	for {
		select {
		case <-stopCh:
			return
			//case 其他协程:

		}
	}
}
