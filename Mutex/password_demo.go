package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM 	sync.RWMutex
	M 		sync.Mutex
	password string
}

var Password = secret{password: "root"}

// 通过rwmutex写
func Change(s *secret, pass string) {
	s.RWM.Lock()
	fmt.Println("Change with rwmutex lock")
	time.Sleep(3 * time.Second)
	s.password = pass
	s.RWM.Unlock()
}

func rwMutexShow(s *secret) string {
	s.RWM.RLock()
	defer s.RWM.RUnlock()
	fmt.Println("show with rwmutex",time.Now().Second())
	time.Sleep(1 * time.Second)
	return s.password
}
// 通过mutex读，和rwMutexShow的唯一区别在于锁的方式不同
func mutexShow(s *secret) string {
	s.M.Lock()
	defer s.M.Unlock()
	fmt.Println("show with mutex",time.Now().Second())
	time.Sleep(1 * time.Second)
	return s.password
}

func main() {
	var show = func(s *secret) string {return ""}

	// 通过变量赋值的方式，选择并重写showFunc函数
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!",time.Now().Second())
		show = rwMutexShow
	} else {
		fmt.Println("Using sync.Mutex!",time.Now().Second())
		show = mutexShow
	}

	var wg sync.WaitGroup

	// 激活5个goroutine，每个goroutine都查看
	// 根据选择的函数不同，showFunc()加锁的方式不同
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", show(&Password),time.Now().Second())
		}()
	}
	// 激活一个申请写锁的goroutine
	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "123456")
	}()

	wg.Wait()
}
