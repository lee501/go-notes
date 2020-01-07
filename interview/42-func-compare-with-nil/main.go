package main

import "fmt"

/*
	考察点：
		func只能跟nil作比较
*/
func funcDemo() {
	var fn1 = func() {}
	var fn2 func()
	if fn1 != nil {
		fmt.Println("fn1 is not nil")
	}
	if fn2 == nil {
		fmt.Println("fn2 is nil")
	}
}

/*
	考察点：
		关于map值为结构体（map[key]struct）
			需要初始化结构体，再赋值，否则struct是不可寻址的
*/
type T struct {
	n int
}
func mapStruct() {
	m := make(map[int]T)

	//初始化struct
	t := T{1}
	m[0] = t
	fmt.Println(m[0].n)
}

func main() {
	funcDemo()
	mapStruct()
}
