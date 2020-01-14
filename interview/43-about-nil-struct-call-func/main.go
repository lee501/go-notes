package main

import "fmt"

/*
	struct方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址
*/
type X struct {

}

func (x *X) callmethod() {
	fmt.Println("test")
}

func main() {
	//nil是合法的调用
	var x *X
	x.callmethod()
	//cannot take the address of X literal, X{}是不可寻址的
	//X{}.callmethod()
	//	正确处理
	a := X{}
	a.callmethod()

	findItemInMap()

	//通过函数创建T
	//GetT().n = 1 无法寻址
	t := GetT()
	t.n = 1 //或 p := &(t.n) *p = 1
	fmt.Println(t)
}

func findItemInMap() {
	x := map[string]string{"one":"a","two":"","three":"c"}
	v := x["one"]
	fmt.Println(v)
	//检查map是否存在值，使用map返回的第二个参数
	if _, ok := x["once"]; !ok {
		fmt.Println("no once")
	}
}

type T struct {
	n int
}

func GetT() T {
	return T{}
}