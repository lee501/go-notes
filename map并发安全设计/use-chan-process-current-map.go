package main

/** * 并发编程，map的线程安全性问题，使用channel的方式 */
import (
	"fmt"
	"time"
)

var dataCh map[int]int = make(map[int]int)
var chMap chan int = make(chan int)
func main() {
	// 并发启动的协程数量
	max := 10000
	time1 := time.Now().UnixNano()

	for i := 0; i < max; i++ {
		go modifyByChan(i)
	}
	// 处理channel的服务
	chanServ(max)
	time2 := time.Now().UnixNano()
	fmt.Printf("data len=%d, time=%d", len(dataCh), (time2-time1)/1000000)
}

func modifyByChan(i int) {
	chMap <- i
}
// 专门处理chMap的服务程序
func chanServ(max int) {
	for {
		i := <- chMap
		dataCh[i] = i
		if len(dataCh) == max {
			return
		}
	}
}