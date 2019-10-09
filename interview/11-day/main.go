package main

import "fmt"

/*
	考点：匿名组合，
		外部类型可以引用内部类型的属性和方法
		外部类型可以定义自己的属性和方法，如果同内部类型相同，则覆盖内部
*/

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
