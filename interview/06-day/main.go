package main

import "fmt"

/*
	考点：
		指针变量p取值方式， 可以通过p.name 或(*p).name, 因为go将p.name转换成（*p).name
*/
type People struct {
	name string
}

func main() {
	p := new(People)
	p.name = "lee"
	fmt.Println(p.name)
	fmt.Println((*p).name)
	convert()
}

/*
	考点：
		类型别名和类型定义
*/
//定义新的类型
type Mint int
//类型别名
type Nint = int

func convert() {
	var i int = 1
	//类型定义需要强制转换
	var j Mint = Mint(i)
	//类型别名可以直接赋值
	var e Nint = i
	fmt.Println(j, e)
}
