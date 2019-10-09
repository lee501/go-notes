package main

import "fmt"

/*
	考点：nil切片和空切片
*/
func nilIssue() {
	//切片是引用类型， 声明指针，定义的是一个nil，没有具体指向值
	var i []int

	//i[0] = 1， 无法直接赋值

	//空切片, 空切片和 nil 不相等，表示一个空的集合
	var j = []int{}
	if j == nil {
		fmt.Println("空切片")
	}

	if i == nil {
		fmt.Println("nil")
	}
}

/*
	考点：一个类型实现多个接口， 向上引用赋值给接口变量
*/

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w *Work) ShowA() int {
	return w.i + 10
}

func (w *Work) ShowB() int {
	return w.i + 20
}

func main() {
	nilIssue()
	w := &Work{5}
	var i A = w
	var j B = w
	fmt.Println(i.ShowA(), j.ShowB())
}
