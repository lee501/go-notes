package main

import (
	"fmt"
	"time"
)

func main() {
	//数组
	var m = [...]int{1,2,3}
	//变量 i、v 是具体元素的副本，在每次循环体中都会被重用， 而不是重新被创建
	for i, v := range m {
		//内部的goroutine 此处可以理解为闭包
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

/*
	1。使用函数传递
		for i, v := range m {
			go func(i,v int) {
				fmt.Println(i, v)
			}(i,v)
		}
	2。重新声明变量
		for i, v := range m {
			i := i           // 这里的 := 会重新声明变量，而不是重用
			v := v
			go func() {
				fmt.Println(i, v)
			}()
		}
*/