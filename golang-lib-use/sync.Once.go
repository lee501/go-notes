package main

/*
	sync.Once是sync包的一个结构体
		type Once struct {
			m    Mutex
			done uint32
		}

		func (o *Once) Do(f func()) {
			if atomic.LoadUint32(&o.done) == 1 {
				return
			}
			// Slow-path.
			o.m.Lock()
			defer o.m.Unlock()
			if o.done == 0 {
				defer atomic.StoreUint32(&o.done, 1)
				f()
			}
		}

	使用场景：
	在程序多次调用的地方只想执行一次某代码块，可以全局声明一个once，使用once.Do()来执行次代码块
*/
import (
	"fmt"
	"sync"
	"time"
)

func demo() {
	var once sync.Once

	for i := 0; i <= 10; i++ {
		go once.Do(func() {
			fmt.Println("hello go")
		})
	}

	time.Sleep(time.Second * 2)
}
