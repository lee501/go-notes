package main

import "fmt"

//interface的methods set规则：*T和T都可是实现接口method的
//值接收者（T）的方法， 结构体指针（*T）和值（T）都可以调用
//指针接收者（*T）的方法， 只有结构体指针（*T）能调用

//接口Valueer
type Valueer interface {
	printf()
}

//value1和value2实现接口Valueer的方法
type Value1 struct {
	num int
}
type Value2 struct {
	num int
}

//接收者为指针
func (self *Value1) printf() {
	fmt.Println(self.num)
}

//接收者为结构体值
func (self Value2) printf() {
	fmt.Println(self.num)
}

func main() {
	//声明接口变量
	var val Valueer
	//Value1赋值给val只能是指针型（*T）
	val = &Value1{num: 1}
	//调用方法
	val.printf()
	//Value2赋值给val可以是值(T)， 也可以是指针（*T）
	val = Value2{num: 2}
	val = &Value2{num: 3}
	//调用方法
	val.printf()
}

