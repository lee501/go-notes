package main

//1. fatal error: all goroutines are asleep - deadlock!

func main() {
	//无缓存通道读
	ch := make(chan int)
	<-ch
	//无缓存通道写，没有写成读
	ch <- 1

	//通道的缓存无数据，但执行读通道。
	//通道的缓存已经占满，向通道写数据，但无协程读。
	chm := make(chan int, 5)
	<- chm
}

//2. panic
//
//向已经关闭的channel写。
//c关闭已经关闭的channel