package main

import "fmt"

/*
1.多态在代码层面最常见的一种方式是: 接口当作方法参数
2.结构体变量赋值给接口变量
3.重写接口时接收者为`Type`和`*Type`的区别
* `*Type`可以调用`*Type`和`Type`作为接收者的方法.所以只要接口中多个方法中至少出现一个使用`*Type`作为接收者进行重写的方法,就必须把结构体指针赋值给接口变量,否则编译报错
* `Type`只能调用`Type`作为接收者的方法
*/
type Live interface {
	run()
}

type People struct {
	name string
}

type Animal struct {

}

func (p *People) run() {
	fmt.Println("人在跑")
}
func (a *Animal) run() {
	fmt.Println("动物在跑")
}
//接口作为参数
func checkSource(live Live) {
	live.run()
}

func main() {
	people := &People{}
	animal := &Animal{}
	//调用方法
	checkSource(people)
	checkSource(animal)
	//将结构体赋值给接口（向上引用）
	var run Live = &People{}
	run.run()
}
