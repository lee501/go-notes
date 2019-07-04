package main

import (
	"fmt"
	"time"
)

//for select中退出的技巧
//使用标签退出
func main()  {
	L:
		for {
			select {
			case <-time.After(time.Second):
					fmt.Println("hello")
			default:
				break L
			}
		}
	fmt.Println("ending")
}

//封装在函数中，使用return退出
func foo() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("hello")
		default:
			return
		}
	}
}

//或者返回一个错误

//if err := foo(); err != nil {
//	panic(err)
//}