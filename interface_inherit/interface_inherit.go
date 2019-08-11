package main

import "fmt"

type Man struct {
	name string
	mile float64
}

type Live interface {
	run(run float64)
	//eat()
	Eat
}

//实现struct man继承接口Live
//重写run和eat方法
func (man *Man) run(mile float64) {
	fmt.Println("跑动了", mile, "公里")
}

func (man *Man) eat() {
	fmt.Println("一天吃两顿烦")
}

func main() {
	man := new(Man)
	man.name = "lee"
	man.mile = 200
	man.run(man.mile)
}

//实现接口的继承
type Eat interface {
	eat()
}

type Child struct {

}

func (child Child) eat() {

}