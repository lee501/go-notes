package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//模拟售票系统
var (
	//票数
	numt = 100
	w sync.WaitGroup
	m sync.Mutex   //互斥锁
)

func main() {
	//设置随机种子，避免每次随机数是一养的
	rand.Seed(time.Now().UnixNano())
	w.Add(4)
	go sellTickert(1)
	go sellTickert(2)
	go sellTickert(3)
	go sellTickert(4)
	w.Wait()
	fmt.Println("所有票数卖完")
	os.Exit(0)
}

func sellTickert(i int) {
	defer func() {
		w.Done()
	}()
	for {
		m.Lock()
		if numt > 0 {
			fmt.Println("第", i, "个窗口卖了", numt)
			numt = numt - 1
		}
		m.Unlock()
		if numt <= 0 {
			break
		}
		//防止一个goroutine卖了所有票
		time.Sleep(time.Duration(rand.Int63n(1000) * 1e6))
	}
}
