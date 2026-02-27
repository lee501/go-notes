package main

import "fmt"

func main() {
	n := f(3)
	fmt.Println(n)
}

func f(n int) (r int) {
	defer func() {
		//此处闭包， n值为参数值
		r += n
		recover()
	}()

	var f func()
	//引发panic
	defer f()

	f = func() {
		r += 2
	}
	return n + 1
}