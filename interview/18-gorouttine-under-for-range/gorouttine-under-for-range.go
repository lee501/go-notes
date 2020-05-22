package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {
	//数组
	var m = [...]int{1,2,3}
	fmt.Println(reflect.TypeOf())
	//变量 i、v 是具体元素的副本，在每次循环体中都会被重用， 而不是重新被创建
	get(m)
	for i, v := range m {
		//内部的goroutine 此处可以理解为闭包
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

func get(a interface{}) {
	fmt.Println(a)
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