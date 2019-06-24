package main

import "fmt"

//方法的结构指针接收者和结构值接收者
//结构指针接收者(*T)和值接收者(T)的方法集是互相继承的
//区别在于：结构体指针（*T）可以改变对象的值

type Cat struct {
	Name string
	Color string
}
//结构体指针类型接收者方法
func (cat *Cat) Meow() {
	fmt.Println("Name:", cat.Name)
}
//结构体值接收者方法
func (cat Cat) Eat() {
	fmt.Println("Color:", cat.Color)
}
func main() {
	//*T和T的方法集是互通的
	//值类型
	cat := Cat{Name: "a"}
	cat.Meow()
	cat.Eat()

	//指针类型
	cater := &Cat{"b", "black"}
	cater.Eat()
	cater.Meow()
}