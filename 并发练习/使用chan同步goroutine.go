package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("执行子线程")
		done <- true
	}()
	<- done
	fmt.Println("执行结束")
}
